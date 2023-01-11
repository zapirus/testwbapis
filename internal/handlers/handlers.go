package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zapirus/testwbapis/internal/service"

	"log"
	"net/http"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Run функция старта
func (s *APIServer) Run() error {
	s.confRouter()
	log.Println("Starting  service")
	return http.ListenAndServe(s.config.HTTPAddr, s.router)
}

// Роуты для запросов
func (s *APIServer) confRouter() {
	// post
	s.router.HandleFunc("/user", s.UserPost())
	s.router.HandleFunc("/shop", s.ShopPost())

	// change
	s.router.HandleFunc("/changeuser/{id}", s.UserChange())
	s.router.HandleFunc("/changeshop/{id}", s.ShopChange())

	//delete
	//s.router.HandleFunc("/")

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

		result, _ := service.UniversalFunc(w, r)
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
		result, _ := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)
	}
}
