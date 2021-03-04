package controller

import (
	"encoding/json"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
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

	tempSatellites := make([]model.TempSatellite, 0)
	helper.OrmConnection.Db.Find(&tempSatellites)

	satelliteData := satellites.Prepare(tempSatellites)

	//shipData := model.ShipDataResponse{
	//	Position: make(map[string]float32),
	//}

	//if x, y, err := location.GetLocation(satelliteData[0].d, distances[1], distances[2]); err != nil {
	//	helper.JsonError(writer, err, http.StatusBadRequest)
	//} else {
	//	shipData.Position["x"] = x
	//	shipData.Position["y"] = y
	//}
	//
	//if solvedMsg, err := messages.GetMessage(msgs[0], msgs[1], msgs[2]); err != nil {
	//	helper.JsonError(writer, err, http.StatusBadRequest)
	//} else {
	//	shipData.Message = solvedMsg
	//}

	res, err := json.Marshal(satelliteData)
	if err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
		return
	}

	helper.JsonOk(writer, res, http.StatusOK)
}
