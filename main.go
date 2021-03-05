package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/router"
)

func main() {
	helper.OrmConnection.InitOrmConnection()
	router.BootstrapServer()
}
