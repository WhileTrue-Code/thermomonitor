package db

import (
	"thermomonitor/model"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type InfluxDB struct {
	Client influxdb2.Client
}

func (influx InfluxDB) Create(thermometer *model.Thermometer){

}

func (influx InfluxDB) Delete(thermometerID uuid.NullUUID){}

func (influx InfluxDB) Update(thermometer *model.Thermometer){}

func (influx InfluxDB) TestCreateThermometer(thermometer model.Thermometer) {
	influxWrite := influx.Client.WriteAPI("","")
	influxWrite.WritePoint().
}
