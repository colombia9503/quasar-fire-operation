package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	helper.OrmConnection.InitOrmConnection()

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
