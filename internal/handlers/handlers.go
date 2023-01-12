package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zapirus/testwbapis/internal/service"
	"net/http"
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

// Run функция старта
func (s *APIServer) Run() error {
	s.confRouter()
	s.logger.Println("Starting server")
	return http.ListenAndServe(s.config.HTTPAddr, s.router)
}

// Роуты для запросов
func (s *APIServer) confRouter() {
	// POST
	s.router.HandleFunc("/user", s.UserPost())
	s.router.HandleFunc("/shop", s.ShopPost())

	// Изменение, или удаление. (в зависимости от запроса)
	s.router.HandleFunc("/changeuser/{id}", s.UserChange())
	s.router.HandleFunc("/changeshop/{id}", s.ShopChange())

	// получение всех записей
	s.router.HandleFunc("/getallusers", s.GetAllUsers()).Methods("GET")
	s.router.HandleFunc("/getallshops", s.GetAllShops()).Methods("GET")

	//роуты для юзера
	s.router.HandleFunc("/getoneuser/{title}", s.GetOneUser())
	s.router.HandleFunc("/getoneshop/{title}", s.GetOneShop())

	s.router.HandleFunc("/getfielduser/{title}", s.GetOneFieldUser()).Methods("GET")
	s.router.HandleFunc("/getfieldshop/{title}", s.GetOneFieldShop()).Methods("GET")

}

// GetAllUsers Функция, которая получает всех юзеров
func (s *APIServer) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := service.GetAll(w, r)
		json.NewEncoder(w).Encode(res)
	}

}

// GetOneUser Функция, которая получает одну запись
func (s *APIServer) GetOneUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := service.GetOneTable(w, r)
		json.NewEncoder(w).Encode(res)
	}

}

// GetAllShops Функция, которая получает все магазины
func (s *APIServer) GetAllShops() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, res := service.GetAll(w, r)
		json.NewEncoder(w).Encode(res)
	}

}

// GetOneShop Функция, которая получает одну запись
func (s *APIServer) GetOneShop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, res := service.GetOneTable(w, r)
		json.NewEncoder(w).Encode(res)
	}

}

// GetOneFieldUser Функция, которая возвращает одно поле по названию
func (s *APIServer) GetOneFieldUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := service.GetOneField(w, r)
		json.NewEncoder(w).Encode(res)

	}

}

// GetOneFieldShop Функция, которая возвращает одно поле по названию
func (s *APIServer) GetOneFieldShop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, res := service.GetOneField(w, r)
		json.NewEncoder(w).Encode(res)
	}

}

// UserPost Функция, которая создает юзеров
func (s *APIServer) UserPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, _ := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)

	}

}

// ShopPost Функция, которая создает магазин
func (s *APIServer) ShopPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, result := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)
	}

}

// UserChange Функция для изменения юзера
func (s *APIServer) UserChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		result, _ := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)
	}
}

// ShopChange Функция для изменения магазина
func (s *APIServer) ShopChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, result := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)
	}
}
