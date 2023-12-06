package coupe

import "task2/internal"

type Coupe struct{}

func (c *Coupe) StartEngine() string {
	return "Starting the Coupe engine."
}

func (c *Coupe) Drive() string {
	return "Coupe is driving."
}

type Factory struct{}

func (f *Factory) CreateVehicle() internal.Vehicle {
	return &Coupe{}
}
