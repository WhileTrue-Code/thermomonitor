package main

import (
	"fmt"
	"thermomonitor/db"
)

func main() {

	var (
		myDb db.DB
	)

	myDb = db.InfluxDB{}
	fmt.Println("Welcome to our ThermoMonitor app.")
}
