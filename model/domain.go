package model

import "time"

type (
	SatelliteData struct {
		Name      string `gorm:"primaryKey"`
		Distance  float32
		Message   string
		CreatedAt time.Time `gorm:"autoCreateTime"`
	}

	TempSatellite struct {
		Name      string `gorm:"primaryKey"`
		X         float32
		Y         float32
		CreatedAt time.Time `gorm:"autoCreateTime"`
	}

	Satellite struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	}

	ShipDataResponse struct {
		Position map[string]float32 `json:"position"`
		Message  string             `json:"message"`
	}
)
