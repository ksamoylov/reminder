package middlewares

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"reminder/config"
	"reminder/internal/app/handlers"
	"reminder/pkg/logger"
)

func Middleware(next handlers.HandlerFn, conf *config.Config) handlers.HandlerFn {
	return func(w http.ResponseWriter, r *http.Request) *handlers.StatusError {
		debug(r, conf.DebugMode)

		err := checkToken(r, "access") // todo

		if err != nil {
			return &handlers.StatusError{Err: err, Code: http.StatusUnauthorized}
		}

		err = checkIsMethodAllowed(r)

		if err != nil {
			return &handlers.StatusError{
				Err:  err,
				Code: http.StatusMethodNotAllowed,
			}
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

func checkToken(r *http.Request, accessToken string) error {
	token := r.Header.Get("Authorization")

	if token != accessToken {
		return errors.New("invalid token passed")
	}

	return nil
}

func checkIsMethodAllowed(r *http.Request) error {
	if urls, ok := handlers.MethodRule[r.Method]; ok {
		for _, url := range urls {
			if url == r.URL.Path {
				return nil
			}
		}
	}

	return errors.New("method not allowed")
}
