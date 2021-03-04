package model

type SatellitesReqBody struct {
	Satellites []Satellite `json:"satellites"`
}

func (srb SatellitesReqBody) Prepare() (distances [3]float32, messages [3][]string) {
	for _, s := range srb.Satellites {
		switch s.Name {
		case "kenobi":
			distances[0] = s.Distance
			messages[0] = s.Message
		case "skywalker":
			distances[1] = s.Distance
			messages[1] = s.Message
		case "sato":
			distances[2] = s.Distance
			messages[2] = s.Message
		}
	}

	return distances, messages
}
