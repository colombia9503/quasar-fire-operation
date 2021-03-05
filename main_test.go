package main

import (
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/router"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	helper.OrmConnection.InitOrmConnection()
	router.BootstrapServer()

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

func TestTopSecretController_TrilaterateShipPosition(t *testing.T) {
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/topsecret",
		strings.NewReader(`{"satellites":[{"name":"kenobi","distance":100,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}`))

	response := executeRequest(req)
	checkResponse(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.AppRouter.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponse(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf("Wrong http status, expected: %d actual: %d", expected, actual)
	}
}
