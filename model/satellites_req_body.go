package model

import "strings"

type SatellitesReqBody struct {
	Satellites []Satellite `json:"satellites"`
}

func (srb SatellitesReqBody) Prepare(satellites []TempSatellite) (satelliteData [3]SatelliteData) {
	for k, v := range satellites {
		for _, s := range srb.Satellites {
			if v.ID == s.Name {
				satelliteData[k] = SatelliteData{
					TempSatelliteID: s.Name,
					Distance:        s.Distance,
					Message:         strings.Join(s.Message, "|"),
				}
			}
		}
	}

	return satelliteData
}
