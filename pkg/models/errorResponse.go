package models

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Message: message})
	if err != nil {
		logrus.Debugf("failed Encode: %v", err)
	}

}
