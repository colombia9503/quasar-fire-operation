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

	log.Println("Creating tables using entities")
	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.TempSatellite{})
	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.SatelliteData{})

	log.Println("Inserting default satellites")
	tempSatellites := []model.TempSatellite{
		{
			ID: "kenobi",
			X:  -500,
			Y:  -200,
		},
		{
			ID: "skywalker",
			X:  100,
			Y:  -100,
		},
		{
			ID: "sato",
			X:  500,
			Y:  100,
		},
	}

	helper.OrmConnection.Db.Create(&tempSatellites)

	log.Println("Initializing mux routers")
	var ar router.AppRouter
	ar.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(ar.Router)

	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", n)
}
