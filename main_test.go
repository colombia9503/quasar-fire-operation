package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	helper.OrmConnection.InitOrmConnection()

	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.TempSatellite{})
	_ = helper.OrmConnection.Db.Migrator().CreateTable(&model.SatelliteData{})

	code := m.Run()
	deleteTables()
	os.Exit(code)
}

func deleteTables() {
	if result := helper.OrmConnection.Db.Exec(`
		DROP TABLE satellite_data;
		DROP TABLE temp_satellites;
	`); result.Error != nil {
		log.Fatal(result.Error)
	}
}
