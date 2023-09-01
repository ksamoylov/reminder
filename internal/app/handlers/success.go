package handlers

import "reminder/internal/app/types"

func CreateSuccessfulResponse(message string) *types.Response {
	return types.NewResponse(true, message)
}
