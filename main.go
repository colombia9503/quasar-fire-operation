package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/router"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	helper.OrmConnection.InitOrmConnection()

	log.Println("Initializing mux routers")

	router.AppRouter.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(router.AppRouter.Router)

	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", n)
}
