package main

// Абстрактные продукты

type Fruit interface {
	GetName() string
	GetColor() string
}

type Vegetable interface {
	GetName() string
	GetColor() string
}

var _ FoodFactory = (*HealthyFoodFactory)(nil)

// FoodFactory Абстрактная фабрика
type FoodFactory interface {
	CreateFruit() Fruit
	CreateVegetable() Vegetable
}

// Конкретные продукты

type Apple struct{}

func (a *Apple) GetName() string {
	return "Яблоко"
}

func (a *Apple) GetColor() string {
	return "Красное"
}

type Carrot struct{}

func (c *Carrot) GetName() string {
	return "Морковка"
}

func (c *Carrot) GetColor() string {
	return "Оранжевая"
}

// Конкретные фабрики

type HealthyFoodFactory struct{}

func (hff *HealthyFoodFactory) CreateFruit() Fruit {
	return &Apple{}
}

func (hff *HealthyFoodFactory) CreateVegetable() Vegetable {
	return &Carrot{}
}

type UnhealthyFoodFactory struct{}

func (uff *UnhealthyFoodFactory) CreateFruit() Fruit {
	return &Apple{}
}

func (uff *UnhealthyFoodFactory) CreateVegetable() Vegetable {
	return &Carrot{}
}

func main() {
	healthyFoodFactory := &HealthyFoodFactory{}
	unhealthyFoodFactory := &UnhealthyFoodFactory{}

	healthyFruit := healthyFoodFactory.CreateFruit()
	healthyVegetable := healthyFoodFactory.CreateVegetable()

	unhealthyFruit := unhealthyFoodFactory.CreateFruit()
	unhealthyVegetable := unhealthyFoodFactory.CreateVegetable()

	println(healthyFruit.GetName())
	println(healthyVegetable.GetName())

	println(unhealthyFruit.GetName())
	println(unhealthyVegetable.GetName())
}
