package controller

import (
	"encoding/json"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/services/location"
	"github.com/colombia9503/quasar-fire-operation/services/messages"
	"net/http"
)

type topSecretController struct{}

var TopSecret = new(topSecretController)

func (ts topSecretController) TrilaterateShipPosition(writer http.ResponseWriter, req *http.Request) {
	var satellites model.SatellitesReqBody
	if err := json.NewDecoder(req.Body).Decode(&satellites); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
		return
	}
	distances, msgs := satellites.Prepare()

	shipData := model.ShipDataResponse{
		Position: make(map[string]float32),
	}

	if x, y, err := location.GetLocation(distances[0], distances[1], distances[2]); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
	} else {
		shipData.Position["x"] = x
		shipData.Position["y"] = y
	}

	if solvedMsg, err := messages.GetMessage(msgs[0], msgs[1], msgs[2]); err != nil {
		helper.JsonError(writer, err, http.StatusBadRequest)
	} else {
		shipData.Message = solvedMsg
	}

	res, err := json.Marshal(shipData)
	if err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
		return
	}

	helper.JsonOk(writer, res, http.StatusOK)
}
