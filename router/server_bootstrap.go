package router

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func BootstrapServer() {
	log.Println("Initializing mux routers")
	AppRouter.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(AppRouter.Router)

	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", n)
}
