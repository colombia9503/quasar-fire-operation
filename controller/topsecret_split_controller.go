package controller

import (
	"encoding/json"
	"github.com/colombia9503/quasar-fire-operation/helper"
	"github.com/colombia9503/quasar-fire-operation/model"
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
		Name:     satellite.Name,
		Message:  strings.Join(satellite.Message, "|"),
		Distance: satellite.Distance,
	}

	if result := helper.OrmConnection.Db.Save(&satelliteData); result.Error != nil {
		helper.JsonError(writer, result.Error, 400)
		return
	}

	if res, err := json.Marshal(satelliteData); err != nil {
		helper.JsonError(writer, err, http.StatusInternalServerError)
	} else {
		helper.JsonOk(writer, res, http.StatusCreated)
	}
}
