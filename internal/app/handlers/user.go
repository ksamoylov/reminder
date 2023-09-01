package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reminder/internal/app/types"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) *types.StatusError {
	user, err := h.deps.UserService.Create(r.Body)

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	response, err := json.Marshal(CreateSuccessfulResponse(fmt.Sprintf("User %d created", user.ID)))

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	w.Write(response)

	return nil
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) *types.StatusError {
	ctx := context.Background()

	token, err := h.deps.UserService.Auth(r.Body, h.deps.Redis, ctx)

	if err != nil {
		return types.NewStatusError(err, http.StatusUnauthorized)
	}

	var response []byte

	response, err = json.Marshal(types.NewAuthTokenResponse(true, *token))

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	w.Write(response)

	return nil
}
