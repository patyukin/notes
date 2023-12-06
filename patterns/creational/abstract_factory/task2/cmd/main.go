package main

import (
	"task2/internal/suv"
)

func main() {
	suvFactory := suv.Factory{}
	vehicle := suvFactory.CreateVehicle()
	vehicle.StartEngine()
	vehicle.Drive()
}
