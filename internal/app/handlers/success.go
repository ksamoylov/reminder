package handlers

import "reminder/internal/app/types"

func CreateSuccessfulResponse() *types.Response {
	return types.NewResponse(true, "successful request")
}
