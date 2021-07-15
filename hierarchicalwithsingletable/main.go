package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kazu1029/dynamodb-go/hierarchicalwithsingletable/sensors"
	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
)

func main() {
	tableName := "SensorsTable"
	ctx := context.Background()
	db, cleanup := dynamo.SetupTable(ctx, tableName, "./hierarchicalwithsingletable/template.yml")
	defer cleanup()

	sensor := sensors.Sensor{
		ID:       "sensor-1",
		City:     "Poznan",
		Building: "A",
		Floor:    "1",
		Room:     "123",
	}

	manager := sensors.NewManager(db, tableName)
	err := manager.Register(ctx, sensor)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n============================ Start Get ===========================\n")
	returned, err := manager.Get(ctx, "sensor-1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("returned: %v\n", returned)
	fmt.Printf("============================ End Get ===========================\n")

	fmt.Printf("\n============================ Start Do not allow to register many times ===========================\n")
	err = manager.Register(ctx, sensor)
	if err != nil {
		fmt.Printf("do not allow to register many times error: %v\n", err)
	}
	fmt.Printf("============================ End Do not allow to register many times ===========================\n")

	fmt.Printf("\n============================ Start Save new reading ===========================\n")
	err = manager.SaveReading(ctx, sensors.Reading{
		SensorID: "sensor-1", Value: "0.67", ReadAt: time.Now(),
	})
	if err != nil {
		panic(err)
	}

	_, latest, err := manager.LatestReadings(ctx, "sensor-1", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("latest: %v\n", latest)
	fmt.Printf("============================ End Save new reading ===========================\n")

	fmt.Printf("\n============================ Start Get last readings and sensor ===========================\n")
	_ = manager.SaveReading(ctx, sensors.Reading{SensorID: "sensor-1", Value: "0.3", ReadAt: time.Now().Add(-20 * time.Second)})
	_ = manager.SaveReading(ctx, sensors.Reading{SensorID: "sensor-1", Value: "0.5", ReadAt: time.Now().Add(-10 * time.Second)})

	lastReadingSensor, latest, err := manager.LatestReadings(ctx, "sensor-1", 2)
	if err != nil {
		panic(err)
	}
	for _, s := range latest {
		fmt.Printf("latest: %v\n", s)
	}
	fmt.Printf("lastReadingSensor: %v\n", lastReadingSensor)
	fmt.Printf("============================ End Get last readings and sensor ===========================\n")

	fmt.Printf("\n============================ Start Get by sensors by location ===========================\n")
	_ = manager.Register(ctx, sensors.Sensor{ID: "sensor-1", City: "Poznan", Building: "A", Floor: "1", Room: "2"})
	_ = manager.Register(ctx, sensors.Sensor{ID: "sensor-2", City: "Poznan", Building: "A", Floor: "2", Room: "4"})
	_ = manager.Register(ctx, sensors.Sensor{ID: "sensor-3", City: "Poznan", Building: "A", Floor: "2", Room: "5"})

	ids, err := manager.GetSensors(ctx, sensors.Location{City: "Poznan", Building: "A", Floor: "2"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ids: %v\n", ids)
	fmt.Printf("============================ End Get sensors by location ===========================\n")
}
