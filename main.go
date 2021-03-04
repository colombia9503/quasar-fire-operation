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
	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.SatelliteData{})
	helper.OrmConnection.Db.Exec(`
		ALTER TABLE SATELLITE_DATA 
		ADD CONSTRAINT SATELLITE_NAME_FK 
		FOREIGN KEY (NAME) 
		REFERENCES TEMP_SATELLITES(NAME)
	`)

	tempSatellites := []model.TempSatellite{
		{
			Name: "kenobi",
			X:    -500,
			Y:    -200,
		},
		{
			Name: "skywalker",
			X:    100,
			Y:    -100,
		},
		{
			Name: "sato",
			X:    500,
			Y:    100,
		},
	}

	helper.OrmConnection.Db.Create(&tempSatellites)

	var ar router.AppRouter
	ar.InitializeRouters()

	n := negroni.Classic()
	n.UseHandler(ar.Router)

	log.Println("listening on port :8080")
	_ = http.ListenAndServe(":8080", n)
}
