package model

type SatellitesReqBody struct {
	Satellites []Satellite `json:"satellites"`
}

type Satellite struct {
	Name     string   `json:"name"`
	Distance string   `json:"distance"`
	Message  []string `json:"message"`
}

type ShipData struct {
	Position map[string]float32 `json:"position"`
	Message  string             `json:"message"`
}
