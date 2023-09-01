package handlers

import (
	"net/http"
	"reminder/internal/app/types"
)

const (
	SignUp = "/sign-up/"
	SignIn = "/sign-in/"
	List   = "/note/list/"
	Create = "/note/create/"
	Delete = "/note/delete/"
)

var MethodRule = map[string][]string{
	http.MethodGet: {
		List,
	},
	http.MethodPost: {
		SignUp,
		Create,
		SignIn,
	},
	http.MethodDelete: {
		Delete,
	},
}

type Router map[string]func(http.ResponseWriter, *http.Request) *types.StatusError

func (h *Handler) NewRouter() *Router {
	return &Router{
		List:   h.List,
		Create: h.Create,
		Delete: h.Delete,
		SignUp: h.SignUp,
		SignIn: h.SignIn,
	}
}
