package errors

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type CustomError struct {
	Message string `json:"message"`
}

// Error responses with custom json error
func Error(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")

	ce := CustomError{
		Message: msg,
	}

	res, err := json.Marshal(ce)
	if err != nil {
		zap.S().Errorw("marshal", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(res) //nolint
}
