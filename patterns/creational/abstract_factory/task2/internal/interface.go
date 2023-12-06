package internal

type Vehicle interface {
	StartEngine() string
	Drive() string
}

type VehicleFactory interface {
	CreateVehicle() Vehicle
}
