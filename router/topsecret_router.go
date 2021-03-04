package router

import (
	"net/http"
)

func (ar *AppRouter) SetTopSecretRouter() {
	ar.Router.HandleFunc("/api/v1/topsecret", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	}).Methods("POST")
}
