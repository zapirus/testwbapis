package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zapirus/testwbapis/internal/model"
	"github.com/zapirus/testwbapis/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type APIServer struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *APIServer) Run() {
	srv := &http.Server{
		Addr:    s.config.HTTPAddr,
		Handler: s.router,
	}

	s.confRouter()
	s.logger.Printf("Завелись на порту %s", s.config.HTTPAddr)
	idConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Fatalln(err)
		}
		close(idConnClosed)
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalln(err)
		}
	}
	<-idConnClosed
	s.logger.Println("Всего доброго!")
}

// Роуты для запросов
func (s *APIServer) confRouter() {
	// роуты юзера
	s.router.HandleFunc("/user", s.UserPost())
	s.router.HandleFunc("/changeuser/{id}", s.UserChange())
	s.router.HandleFunc("/getallusers", s.GetAllUsers())
	s.router.HandleFunc("/getoneuser/{title}", s.GetOneUser()).Methods("GET")
	s.router.HandleFunc("/getfielduser/{title}", s.GetOneFieldUser()).Methods("GET")

	// роуты шопа
	s.router.HandleFunc("/shop", s.ShopPost())
	s.router.HandleFunc("/changeshop/{id}", s.ShopChange())
	s.router.HandleFunc("/getallshops", s.GetAllShops())
	s.router.HandleFunc("/getoneshop/{title}", s.GetOneShop()).Methods("GET")
	s.router.HandleFunc("/getfieldshop/{title}", s.GetOneFieldShop()).Methods("GET")

}

// Strip Функция, которая режет URL
func (s *APIServer) Strip(url string) string {
	var (
		beforeSpace int
		res         string
	)

	for _, elem := range url {
		if beforeSpace == 2 {
			break
		} else if elem == 47 {
			beforeSpace += 1
		}
		res += string(elem)
	}
	return res
}

// GetAllUsers Функция, которая получает всех юзеров
func (s *APIServer) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		url := r.URL.RequestURI()
		met := r.Method
		res, _ := service.UniversalFunc(met, url, "", model.User{}, model.Shop{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Printf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}

// GetAllShops Функция, которая получает все магазины
func (s *APIServer) GetAllShops() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		url := r.URL.RequestURI()
		met := r.Method
		_, res := service.UniversalFunc(met, url, "", model.User{}, model.Shop{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Printf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}

// GetOneUser Функция, которая получает одну запись
func (s *APIServer) GetOneUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		url := s.Strip(r.URL.RequestURI())
		var reqId = mux.Vars(r)["title"]
		res, _ := service.GetOneEntity(url, reqId)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Printf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}

// GetOneShop Функция, которая получает одну запись
func (s *APIServer) GetOneShop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		url := s.Strip(r.URL.RequestURI())
		var reqId = mux.Vars(r)["title"]
		_, res := service.GetOneEntity(url, reqId)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Errorf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}

// GetOneFieldUser Функция, которая возвращает одно поле по названию
func (s *APIServer) GetOneFieldUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		urlField := s.Strip(r.URL.RequestURI())
		var reqId = mux.Vars(r)["title"]
		res, _ := service.GetOneField(urlField, reqId)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Errorf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}

	}

}

// GetOneFieldShop Функция, которая возвращает одно поле по названию
func (s *APIServer) GetOneFieldShop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		urlField := s.Strip(r.URL.RequestURI())
		var reqId = mux.Vars(r)["title"]
		_, res := service.GetOneField(urlField, reqId)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			s.logger.Printf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}

// UserPost Функция, которая создает юзеров
func (s *APIServer) UserPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" && r.URL.RequestURI() == "/user" {
			met := r.Method
			var newUser model.User
			if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}
			result, _ := service.UniversalFunc(met, r.URL.RequestURI(), "", newUser, model.Shop{})
			if err := json.NewEncoder(w).Encode(result); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}

		} else {
			w.Write([]byte("Что-то пошло не так"))
		}

	}

}

// ShopPost Функция, которая создает магазин
func (s *APIServer) ShopPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" && r.URL.RequestURI() == "/shop" {
			var newShop model.Shop
			met := r.Method
			if err := json.NewDecoder(r.Body).Decode(&newShop); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}
			_, result := service.UniversalFunc(met, r.URL.RequestURI(), "", model.User{}, newShop)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}

		} else {
			w.Write([]byte("Что-то пошло не так"))
		}
	}

}

// UserChange Функция для изменения юзера
func (s *APIServer) UserChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "PUT" && s.Strip(r.URL.RequestURI()) == "/changeuser/" {
			var newUser model.User
			if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}
			url := s.Strip(r.URL.RequestURI())
			met := r.Method
			var reqId = mux.Vars(r)["id"]
			result, _ := service.UniversalFunc(met, url, reqId, newUser, model.Shop{})
			if err := json.NewEncoder(w).Encode(result); err != nil {
				s.logger.Println("Не удалось преобразовать", err)
				w.WriteHeader(http.StatusConflict)
				return
			}

		} else if r.Method == "DELETE" && s.Strip(r.URL.RequestURI()) == "/changeuser/" {
			var newUser model.User
			if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
				s.logger.Println("Не удалось преобразовать", err)
				w.WriteHeader(http.StatusConflict)
				return

			}
			url := s.Strip(r.URL.RequestURI())
			met := r.Method
			var reqId = mux.Vars(r)["id"]
			result, _ := service.UniversalFunc(met, url, reqId, newUser, model.Shop{})
			if err := json.NewEncoder(w).Encode(result); err != nil {
				s.logger.Println("Не удалось преобразовать", err)
				return
			}
		} else {
			s.logger.Println("Ничего не нашлось")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// ShopChange Функция для изменения магазина
func (s *APIServer) ShopChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "PUT" && s.Strip(r.URL.RequestURI()) == "/changeshop/" {
			var newShop model.Shop
			if err := json.NewDecoder(r.Body).Decode(&newShop); err != nil {
				logrus.Fatalln(err)
			}
			url := s.Strip(r.URL.RequestURI())
			met := r.Method
			var reqId = mux.Vars(r)["id"]
			_, result := service.UniversalFunc(met, url, reqId, model.User{}, newShop)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				logrus.Fatalln(err)
			}

		} else if r.Method == "DELETE" && s.Strip(r.URL.RequestURI()) == "/changeshop/" {
			var newShop model.Shop
			if err := json.NewDecoder(r.Body).Decode(&newShop); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}
			url := s.Strip(r.URL.RequestURI())
			met := r.Method
			var reqId = mux.Vars(r)["id"]
			_, result := service.UniversalFunc(met, url, reqId, model.User{}, newShop)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				s.logger.Printf("Не удалось преаброзовать: %s", err)
				w.WriteHeader(http.StatusConflict)
				return
			}
		} else {
			s.logger.Println("Ничего не нашлось")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
