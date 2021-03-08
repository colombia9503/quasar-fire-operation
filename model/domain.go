package model

import "time"

type (
	SatelliteData struct {
		TempSatelliteID string    `json:"name" gorm:"primaryKey"`
		Distance        float64   `json:"distance"`
		Message         string    `json:"message"`
		CreatedAt       time.Time `json:"created_at" gorm:"autoUpdateTime"`
	}

	TempSatellite struct {
		ID            string         `json:"name" gorm:"primaryKey"`
		X             float64        `json:"x"`
		Y             float64        `json:"y"`
		CreatedAt     time.Time      `json:"created_at" gorm:"autoUpdateTime"`
		SatelliteData *SatelliteData `json:"satellite_data" gorm:"foreignKey:TempSatelliteID"`
	}

	Satellite struct {
		Name     string   `json:"name" validate:"notnull,notblank"`
		Distance float64  `json:"distance" validate:"notnull"`
		Message  []string `json:"message" validate:"notnull,nonempty"`
	}

	ShipDataResponse struct {
		Position map[string]float64 `json:"position"`
		Message  string             `json:"message"`
	}
)
