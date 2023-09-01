package app

import (
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"reminder/config"
	"reminder/internal/app/handlers"
	"reminder/internal/app/middlewares"
)

type Server struct {
	config  *config.Config
	handler *handlers.Handler
}

func NewServer(config *config.Config, db *sql.DB, redis *redis.Client) *Server {
	return &Server{config: config, handler: handlers.NewHandler(config, db, redis)}
}

func (s *Server) Start() {
	s.handle()
	s.listen()
}

func (s *Server) handle() {
	router := *s.handler.NewRouter()

	for route, handler := range router {
		http.Handle(route, middlewares.NoteMiddleware(handler, s.config))
	}
}

func (s *Server) listen() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.HttpPort), nil)

	if err != nil {
		fmt.Println(err)
	}
}
