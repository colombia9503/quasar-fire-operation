package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var OrmConnection = new(gormDb)

type gormDb struct {
	Db *gorm.DB
}

func (gdb *gormDb) InitOrmConnection() {

	if port, err := strconv.Atoi(os.Getenv("APP_DB_PORT")); err == nil {
		dbConnectionInstance.BootstrapDBConnection(
			os.Getenv("APP_DB_USERNAME"),
			os.Getenv("APP_DB_PASSWORD"),
			os.Getenv("APP_DB_NAME"),
			os.Getenv("APP_DB_HOST"),
			port,
		)
	}

	var err error
	if gdb.Db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: dbConnectionInstance.Db,
	}), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}
}
