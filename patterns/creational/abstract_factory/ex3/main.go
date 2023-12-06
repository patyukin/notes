package main

import "fmt"

var _ TransportFactory = (*LogisticsFactory)(nil)

// TransportFactory Абстрактный интерфейс фабрики транспорта
type TransportFactory interface {
	CreateGroundTransport() GroundTransport
	CreateAirTransport() AirTransport
}

// GroundTransport Абстрактный интерфейс земного транспорта
type GroundTransport interface {
	Deliver()
}

// AirTransport Абстрактный интерфейс воздушного транспорта
type AirTransport interface {
	Deliver()
}

// LogisticsFactory Реализация фабрики транспорта для доставки при сотрудничестве с земными и воздушными перевозчиками
type LogisticsFactory struct{}

func (lf LogisticsFactory) CreateGroundTransport() GroundTransport {
	return &Truck{}
}

func (lf LogisticsFactory) CreateAirTransport() AirTransport {
	return &Airplane{}
}

// Truck Реализация земного транспорта - грузовика
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Грузовик доставляет груз по дорогам.")
}

// Airplane Реализация воздушного транспорта - самолета
type Airplane struct{}

func (a *Airplane) Deliver() {
	fmt.Println("Самолет доставляет груз по воздуху.")
}

func main() {
	// Создание фабрики для доставки товаров
	factory := LogisticsFactory{}

	// Создание земного транспорта
	groundTransport := factory.CreateGroundTransport()
	groundTransport.Deliver()

	// Создание воздушного транспорта
	airTransport := factory.CreateAirTransport()
	airTransport.Deliver()
}

// В этом примере мы используем шаблон проектирования "Абстрактная фабрика" для создания разных типов транспорта для доставки товаров.
// У нас есть абстрактный интерфейс `TransportFactory`, который определяет методы `CreateGroundTransport()`
// и `CreateAirTransport()` для создания объектов земного и воздушного транспорта соответственно.
// Абстрактные интерфейсы `GroundTransport` и `Airtransport` определяют метод `Deliver()` для доставки грузов.
// Реализация `LogisticsFactory` представляет фабрику транспорта, которая создает конкретные земные транспорты типа `Truck` и воздушные транспорты типа `Airplane`.
// В функции `main` мы создаем экземпляр фабрики `LogisticsFactory` и используем ее для создания объектов земного и воздушного транспорта.
// Затем мы вызываем метод `Deliver()` для каждого типа транспорта, чтобы показать, как каждый тип доставляет груз.
// Этот пример иллюстрирует, как шаблон "Абстрактная фабрика" может быть использован для создания связанных объектов различных типов,
// в нашем случае, земного и воздушного транспорта для доставки товаров.
