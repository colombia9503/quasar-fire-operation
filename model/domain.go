package model

import "time"

type (
	TempSatellite struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		Distance  float32
		Message   string
		BatchNum  int       `gorm:"index"`
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
