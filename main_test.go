package main

import (
	"bytes"
	"encoding/json"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
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
	shipDataResponse := &model.ShipDataResponse{}
	readBody(shipDataResponse, t, response.Body)
	if shipDataResponse.Message != "este es un mensaje secreto" {
		t.Fatalf("Assertion failed: %s", shipDataResponse.Message)
	}
}

func TestTopSecretController_TrilaterateShipPositionBadReq(t *testing.T) {
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/topsecret",
		strings.NewReader(`{"satellites":[{"name":"kenobi","distance":100,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]}]}`))

	response := executeRequest(req)
	checkResponse(t, http.StatusBadRequest, response.Code)
	appErr := &helper.AppErr{}
	readBody(appErr, t, response.Body)

	if appErr.Status != 400 {
		t.Fatalf("Expected 400, actual: %d", appErr.Status)
	} else {
		t.Log(appErr)
	}
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

func readBody(i interface{}, t *testing.T, resBody *bytes.Buffer) {
	if err := json.NewDecoder(resBody).Decode(&i); err != nil {
		t.Fatalf("cannot unmarshal json into given interface")
	} else {
		t.Log(i)
	}
}
