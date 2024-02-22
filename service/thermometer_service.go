package service

import (
	"thermomonitor/db"
	"thermomonitor/model"

	"github.com/google/uuid"
)

type Service struct {
	db db.DB
	//logger
}

func CreateThermometer(thermometer *model.Thermometer) {
	thermometer.ID = uuid.NullUUID{
		UUID:  uuid.New(),
		Valid: false,
	}
}

func DeleteThermometer(thermometerID uuid.NullUUID) {
}

func UpdateThermometer(thermometer *model.Thermometer) {
}
