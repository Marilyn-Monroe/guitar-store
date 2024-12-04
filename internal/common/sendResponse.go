package common

import (
	"encoding/json"
	"guitarStore/internal/model"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, message string) {
	errorMessage := model.ResponseModel{Message: &message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(errorMessage)
}
