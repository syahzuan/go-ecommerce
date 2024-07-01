package api

import (
	"database/sql"
	"go-ecommerce/services/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	userstore := user.NewStore(s.db)
	userHandler := user.NewHandler(userstore)
	apiMux.HandleFunc("POST /login", userHandler.Login)
	apiMux.HandleFunc("POST /register", userHandler.Register)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1/", apiMux))

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, nil)
}
