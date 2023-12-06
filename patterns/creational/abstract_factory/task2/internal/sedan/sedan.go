package sedan

import "task2/internal"

type Sedan struct{}

func (s *Sedan) StartEngine() string {
	return "Starting the Sedan engine."
}

func (s *Sedan) Drive() string {
	return "Sedan is driving."
}

type Factory struct{}

func (s *Factory) CreateVehicle() internal.Vehicle {
	return &Sedan{}
}
