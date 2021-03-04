package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/router"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	helper.OrmConnection.InitOrmConnection()
	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.TempSatellite{})

	var ar router.AppRouter
	ar.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(ar.Router)

	log.Println("listening on port :8081")
	_ = http.ListenAndServe(":8080", n)
}
