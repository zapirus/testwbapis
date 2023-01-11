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

// Start функция старта
func (s *APIServer) Start() error {
	s.confRouter()
	log.Println("Starting  service")
	return http.ListenAndServe(s.config.HTTPAddr, s.router)
}

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/user", s.User())
	s.router.HandleFunc("/shop", s.Shop())
}

func (s *APIServer) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, _ := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)

	}

}

func (s *APIServer) Shop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, _ := service.UniversalFunc(w, r)
		json.NewEncoder(w).Encode(result)
	}

}
