package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	AppError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	objErr := &AppError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(objErr); err == nil {
		w.Write(j)
	}
}
