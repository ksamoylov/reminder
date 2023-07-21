package handlers

import (
	"encoding/json"
	"net/http"
	"reminder/internal/app/types"
	"reminder/pkg/logger"
)

func handleError(w http.ResponseWriter, statusError *StatusError) {
	w.WriteHeader(statusError.Code)
	w.Write(CreateResponseByError(statusError.Err))
}

func CreateResponseByError(err error) []byte {
	response := types.NewResponse(false, err.Error())
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		logger.Error(err)
		return nil
	}

	return jsonResponse
}
