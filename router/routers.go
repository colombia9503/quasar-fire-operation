package router

import "github.com/gorilla/mux"

var AppRouter = new(appRouter)

type appRouter struct {
	Router *mux.Router
}

func (ar *appRouter) InitializeRouters() {
	ar.Router = mux.NewRouter().StrictSlash(false)
	ar.SetTopSecretRouter()
}
