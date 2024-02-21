package db

import (
	"thermomonitor/model"

	"github.com/google/uuid"
)

type DB interface {
	Create(thermometer *model.Thermometer)
	Delete(thermometerID uuid.NullUUID)
	Update(thermometer *model.Thermometer)
}
