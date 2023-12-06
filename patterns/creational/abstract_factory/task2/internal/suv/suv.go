package suv

import "task2/internal"

type SUV struct{}

func (s SUV) StartEngine() string {
	return "Starting the SUV engine."
}

func (s SUV) Drive() string {
	return "SUV is driving."
}

type Factory struct{}

func (suv *Factory) CreateVehicle() internal.Vehicle {
	return &SUV{}
}
