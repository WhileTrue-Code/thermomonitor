package db

import (
	"fmt"
	"thermomonitor/model"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

const (
	thermometerMeasurement      = "thermometer"
	thermometerStateMeasurement = "thermometer_state"
)

type InfluxDB struct {
	Client   influxdb2.Client
	WriteAPI api.WriteAPI
}

func newInfluxDB(host, port string) InfluxDB {
	dbUrl := fmt.Sprintf("%s:%s", host, port)
	client := influxdb2.NewClient(dbUrl, "")
	writeAPI := client.WriteAPI("thermomonitor", "")

	return InfluxDB{
		Client:   client,
		WriteAPI: writeAPI,
	}

}

func (influx InfluxDB) Create(thermometer *model.Thermometer) {
	tags := map[string]string{"id": thermometer.ID.UUID.String()}
	fields := map[string]interface{}{"name": thermometer.Name, "deleted_at": thermometer.DeletedAt}
	point := influxdb2.NewPoint(thermometerMeasurement, tags, fields, time.Now())
	influx.WriteAPI.WritePoint(point)
}

func (influx InfluxDB) Delete(thermometerID uuid.NullUUID) {
}

func (influx InfluxDB) Update(thermometer *model.Thermometer) {}

/*
func (influx InfluxDB) TestCreateThermometer(thermometer model.Thermometer) {
	influxWrite := influx.Client.WriteAPI("","")
	influxWrite.WritePoint().
}
*/
