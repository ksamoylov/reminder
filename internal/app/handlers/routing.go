package handlers

import "net/http"

const (
	List   = "/note/list/"
	Create = "/note/create/"
	Delete = "/note/delete/"
)

var MethodRule = map[string][]string{
	http.MethodGet: {
		List,
	},
	http.MethodPost: {
		Create,
	},
	http.MethodDelete: {
		Delete,
	},
}

type Router map[string]func(http.ResponseWriter, *http.Request) *StatusError

func (h *Handler) NewRouter() *Router {
	return &Router{
		List:   h.List,
		Create: h.Create,
		Delete: h.Delete,
	}
}
