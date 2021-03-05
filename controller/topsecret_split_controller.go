package controller

import (
	"encoding/json"
	"errors"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/services/location"
	"github.com/colombia9503/quasar-fire-operation/services/messages"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type topSecretSplitController struct{}

var TopSecretSplit = new(topSecretSplitController)

func (tss topSecretSplitController) SaveSatelliteData(writer http.ResponseWriter, req *http.Request) {
	var satellite model.Satellite

	satelliteName := mux.Vars(req)["name"]
	if err := json.NewDecoder(req.Body).Decode(&satellite); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	}

	satellite.Name = satelliteName

	satelliteData := model.SatelliteData{
		ID:       satellite.Name,
		Message:  strings.Join(satellite.Message, "|"),
		Distance: satellite.Distance,
	}

	if result := helper.OrmConnection.Db.Save(&satelliteData); result.Error != nil {
		helper.JsonError(writer, errors.New("satellite not found"), 404)
		return
	}

	if res, err := json.Marshal(satelliteData); err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
	} else {
		helper.JsonOk(writer, res, http.StatusCreated)
	}
}

func (tss topSecretSplitController) TrilaterateShipPositionAndMessage(writer http.ResponseWriter, req *http.Request) {
	tempSatellites := make([]model.TempSatellite, 0)
	if result := helper.OrmConnection.Db.Preload("SatelliteData").Find(&tempSatellites); result.Error != nil {
		helper.JsonError(writer, result.Error, http.StatusInternalServerError)
		return
	}

	satellitesData := [3]model.SatelliteData{}
	for k, v := range tempSatellites {
		if v.SatelliteData == nil {
			helper.JsonError(writer, errors.New("the ship message and position cannot be determined"), http.StatusNotFound)
			return
		} else {
			satellitesData[k] = *v.SatelliteData
		}
	}

	shipData := model.ShipDataResponse{
		Position: make(map[string]float32),
	}

	if x, y, err := location.GetShipLocation(tempSatellites, satellitesData); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	} else {
		shipData.Position["x"] = x
		shipData.Position["y"] = y
	}

	msgs := make([][]string, 0)
	for _, v := range satellitesData {
		msgs = append(msgs, strings.Split(v.Message, "|"))
	}

	if solvedMsg, err := messages.GetMessage(msgs[0], msgs[1], msgs[2]); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	} else {
		shipData.Message = solvedMsg
	}

	helper.OrmConnection.Db.Where("1 = 1").Delete(&model.SatelliteData{})

	if res, err := json.Marshal(shipData); err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
	} else {
		helper.JsonOk(writer, res, http.StatusOK)
	}
}
