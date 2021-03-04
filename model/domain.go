package model

import "time"

type (
	SatelliteData struct {
		Name      string    `gorm:"primaryKey" json:"name"`
		Distance  float32   `json:"distance"`
		Message   string    `json:"message"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
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
