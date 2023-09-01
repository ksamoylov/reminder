package handlers

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gopkg.in/validator.v2"
	"net/http"
	"reminder/config"
	"reminder/internal/app/providers"
	"reminder/internal/app/types"
)

type Handler struct {
	config    *config.Config
	validator *validator.Validator
	deps      *providers.Deps
}

func NewHandler(config *config.Config, db *sql.DB, redis *redis.Client) *Handler {
	return &Handler{config: config, deps: providers.NewDeps(db, redis)}
}

type HandlerFn func(w http.ResponseWriter, r *http.Request) *types.StatusError

func (fn HandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setDefaultResponseHeaders(w)
	statusError := fn(w, r)

	if statusError != nil {
		handleError(w, statusError)
	}
}

func setDefaultResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
