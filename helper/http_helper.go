package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	AppErr struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}
	logErr struct {
		Message string `json:"message"`
	}
)

func JsonError(w http.ResponseWriter, handlerErr error, code int) {
	log.Printf("Error: %s\n", handlerErr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if res, err := json.Marshal(&AppErr{handlerErr.Error(), code}); err == nil {
		w.Write(res)
	}
}

func JsonOk(w http.ResponseWriter, res []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(res)
}

func JsonStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func (e *logErr) Error() string {
	return e.Message
}

func NewLogErr(msg string) error {
	return &logErr{msg}
}
