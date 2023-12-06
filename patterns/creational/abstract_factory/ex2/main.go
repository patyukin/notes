package main

import "fmt"

var _ AbstractVehicleFactory = (*FordFactory)(nil)
var _ AbstractVehicleFactory = (*ChevroletFactory)(nil)

// AbstractVehicleFactory интерфейс абстрактной фабрики
type AbstractVehicleFactory interface {
	CreateCar() Car
	CreateMotorcycle() Motorcycle
}

// Car интерфейс автомобиля
type Car interface {
	Drive() string
}

// Motorcycle интерфейс мотоцикла
type Motorcycle interface {
	Ride() string
}

// FordFactory конкретная фабрика для создания автомобилей и мотоциклов от Ford
type FordFactory struct{}

func (f *FordFactory) CreateCar() Car {
	return &FordCar{}
}

func (f *FordFactory) CreateMotorcycle() Motorcycle {
	return &FordMotorcycle{}
}

// ChevroletFactory конкретная фабрика для создания автомобилей и мотоциклов от Chevrolet
type ChevroletFactory struct{}

func (f *ChevroletFactory) CreateCar() Car {
	return &ChevroletCar{}
}

func (f *ChevroletFactory) CreateMotorcycle() Motorcycle {
	return &ChevroletMotorcycle{}
}

// FordCar конкретная реализация автомобиля от Ford
type FordCar struct{}

func (c *FordCar) Drive() string {
	return "Driving a Ford car."
}

// FordMotorcycle конкретная реализация мотоцикла от Ford
type FordMotorcycle struct{}

func (m *FordMotorcycle) Ride() string {
	return "Riding a Ford motorcycle."
}

// ChevroletCar конкретная реализация автомобиля от Chevrolet
type ChevroletCar struct{}

func (c *ChevroletCar) Drive() string {
	return "Driving a Chevrolet car."
}

// ChevroletMotorcycle конкретная реализация мотоцикла от Chevrolet
type ChevroletMotorcycle struct{}

func (m *ChevroletMotorcycle) Ride() string {
	return "Riding a Chevrolet motorcycle."
}

// В этом примере мы имитируем создание автомобилей и мотоциклов от двух разных производителей: Ford и Chevrolet.
// В качестве абстрактной фабрики выступает интерфейс `AbstractVehicleFactory` с методами `CreateCar()` и `CreateMotorcycle()`,
// которые должны быть реализованы конкретными фабриками.
// У нас есть две конкретные фабрики: `FordFactory` и `ChevroletFactory`, каждая из которых реализует интерфейс `AbstractVehicleFactory`.
// При вызове методов создания автомобилей и мотоциклов фабрики возвращают соответствующие объекты для конкретного производителя.
// Каждый тип автомобиля и мотоцикла имеет свою реализацию метода `Drive()` или `Ride()`, который выводит сообщение о том,
// что мы ездим на соответствующем транспортном средстве.
// В функции `main` мы создаем экземпляры фабрик для Ford и Chevrolet, а затем используем их для создания автомобилей и мотоциклов.
// Мы вызываем методы `Drive()` и `Ride()` у каждого транспортного средства и выводим результаты на экран.
// Этот пример демонстрирует, как абстрактная фабрика может использоваться для создания различных типов объектов в зависимости
// от конкретного производителя или семейства продуктов.

func main() {
	// Создание фабрики для автомобилей и мотоциклов от Ford
	fordFactory := FordFactory{}
	fordCar := fordFactory.CreateCar()
	fordMotorcycle := fordFactory.CreateMotorcycle()
	fmt.Println(fordCar.Drive())       // Output: Driving a Ford car.
	fmt.Println(fordMotorcycle.Ride()) // Output: Riding a Ford motorcycle.

	// Создание фабрики для автомобилей и мотоциклов от Chevrolet
	chevroletFactory := ChevroletFactory{}
	chevroletCar := chevroletFactory.CreateCar()
	chevroletMotorcycle := chevroletFactory.CreateMotorcycle()
	fmt.Println(chevroletCar.Drive())       // Output: Driving a Chevrolet car.
	fmt.Println(chevroletMotorcycle.Ride()) // Output: Riding a Chevrolet motorcycle.
}
