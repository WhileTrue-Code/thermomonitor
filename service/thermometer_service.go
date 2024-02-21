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

}

func DeleteThermometer(thermometerID uuid.NullUUID) {
}

func UpdateThermometer(thermometer *model.Thermometer) {
}
