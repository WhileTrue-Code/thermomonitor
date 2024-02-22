package model

import "github.com/google/uuid"

type Thermometer struct {
	ID        uuid.NullUUID `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	DeletedAt *int64        `json:"deletedAt,omitempty"`
}

type ThermometerState struct {
	ThermometerID uuid.NullUUID `json:"thermometerID,omitempty"`
	Temperature   float64       `json:"temperature,omitempty"`
	PH            float64       `json:"ph,omitempty"`
	Humidity      float64       `json:"humidity,omitempty"`
	Timestamp     int64         `json:"timestamp,omitempty"`
}
