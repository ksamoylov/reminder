package handlers

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gopkg.in/validator.v2"
	"net/http"
	"reminder/config"
	"reminder/internal/app/providers"
	"reminder/internal/app/types"
)

const UserIdContextKey = "user_id"

type Handler struct {
	config    *config.Config
	validator *validator.Validator
	Deps      *providers.Deps
}

func NewHandler(config *config.Config, db *sql.DB, redis *redis.Client) *Handler {
	return &Handler{config: config, Deps: providers.NewDeps(db, redis)}
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

func PassUserIdToRequestContext(r *http.Request, ctx context.Context, userId int) {
	req := r.WithContext(context.WithValue(ctx, UserIdContextKey, userId))
	*r = *req
}

func getUserIdFromRequestContext(r *http.Request) int {
	return r.Context().Value("user_id").(int)
}
