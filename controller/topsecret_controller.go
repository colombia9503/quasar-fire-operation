package controller

import (
	"encoding/json"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/services/location"
	"github.com/colombia9503/quasar-fire-operation/services/messages"
	"log"
	"net/http"
	"strings"
)

type topSecretController struct{}

var TopSecret = new(topSecretController)

func (ts topSecretController) TrilaterateShipPosition(writer http.ResponseWriter, req *http.Request) {
	var satellites model.SatellitesReqBody
	if err := json.NewDecoder(req.Body).Decode(&satellites); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	}

	tempSatellites := make([]model.TempSatellite, 0)
	helper.OrmConnection.Db.Find(&tempSatellites)

	satelliteData := satellites.Prepare(tempSatellites)

	shipData := model.ShipDataResponse{
		Position: make(map[string]float64),
	}

	if x, y, err := location.GetShipLocation(tempSatellites, satelliteData); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	} else {
		log.Println("ship location: ", x, y)
		shipData.Position["x"] = x
		shipData.Position["y"] = y
	}

	msgs := make([][]string, 0)
	for _, v := range satelliteData {
		msgs = append(msgs, strings.Split(v.Message, "|"))
	}

	if solvedMsg, err := messages.GetMessage(msgs[0], msgs[1], msgs[2]); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	} else {
		shipData.Message = solvedMsg
	}

	if res, err := json.Marshal(shipData); err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
	} else {
		helper.JsonOk(writer, res, http.StatusOK)
	}
}
