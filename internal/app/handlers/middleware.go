package handlers

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"net/http/httputil"
	"reminder/config"
	"reminder/internal/app/types"
	"reminder/internal/app/util"
	"reminder/pkg/logger"
)

const authorizationHeader = "Authorization"

func CommonMiddleware(next HandlerFn, conf *config.Config, redis *redis.Client) HandlerFn {
	return func(w http.ResponseWriter, r *http.Request) *types.StatusError {
		var err error
		ctx := r.Context()

		debug(r, conf.DebugMode)

		err = checkIsMethodAllowed(r)

		if err != nil {
			return types.NewStatusError(err, http.StatusMethodNotAllowed)
		}

		err = checkAuth(r, redis, ctx)

		if err != nil {
			return types.NewStatusError(err, http.StatusUnauthorized)
		}

		return next(w, r)
	}
}

func debug(r *http.Request, debugMode bool) {
	if !debugMode {
		return
	}

	out, err := httputil.DumpRequest(r, true)

	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(string(out))
}

func checkToken(r *http.Request, redis *redis.Client, ctx context.Context) error {
	token := r.Header.Get(authorizationHeader)

	tokenClaims, err := util.ParseToken(token)

	if err != nil {
		return err
	}

	isSessionExists := util.CheckSessionOnExists(tokenClaims, redis, ctx)

	if !isSessionExists {
		return errors.New("no active session")
	}

	PassUserIdToRequestContext(r, ctx, tokenClaims.UserId)

	return nil
}

func checkAuth(r *http.Request, redis *redis.Client, ctx context.Context) error {
	if !checkIsMethodNeedAuth(r) {
		return nil
	}

	err := checkToken(r, redis, ctx)

	if err != nil {
		return err
	}

	return nil
}

func checkIsMethodNeedAuth(r *http.Request) bool {
	for _, url := range MethodsAuthNeeded {
		if url == r.URL.Path {
			return true
		}
	}

	return false
}

func checkIsMethodAllowed(r *http.Request) error {
	if urls, ok := MethodRule[r.Method]; ok {
		for _, url := range urls {
			if url == r.URL.Path {
				return nil
			}
		}
	}

	return errors.New("method not allowed")
}
