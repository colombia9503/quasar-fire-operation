package router

import (
	"github.com/colombia9503/quasar-fire-operation/controller"
)

func (ar *AppRouter) SetTopSecretRouter() {
	ar.Router.HandleFunc("/api/v1/topsecret", controller.TopSecret.TrilaterateShipPosition).Methods("POST")
	ar.Router.HandleFunc("/api/v1/topsecret_split/{name}", controller.TopSecretSplit.SaveSatelliteData).Methods("POST")
	ar.Router.HandleFunc("/api/v1/topsecret_split", controller.TopSecretSplit.TrilaterateShipPositionAndMessage).Methods("GET")
}
