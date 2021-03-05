package model

import "time"

type (
	SatelliteData struct {
		ID        string    `json:"name" gorm:"primaryKey"`
		Distance  float32   `json:"distance"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime"`
	}

	TempSatellite struct {
		ID            string         `json:"name" gorm:"primaryKey"`
		X             float32        `json:"x"`
		Y             float32        `json:"y"`
		CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
		SatelliteData *SatelliteData `json:"satellite_data" gorm:"foreignKey:ID"`
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
