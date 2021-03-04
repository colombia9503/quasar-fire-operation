package router

import "github.com/gorilla/mux"

type AppRouter struct {
	Router *mux.Router
}

func (ar *AppRouter) InitializeRouters() {
	ar.Router = mux.NewRouter().StrictSlash(false)
	ar.SetTopSecretRouter()
}
