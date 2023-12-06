package main

import "fmt"

// ProductFactory Абстрактный интерфейс фабрики продуктов
type ProductFactory interface {
	CreateProduct() Product
	CreateShippingService() ShippingService
}

// Product Абстрактный интерфейс продукта
type Product interface {
	GetInfo() string
}

// ShippingService Абстрактный интерфейс службы доставки
type ShippingService interface {
	Deliver()
}

// ElectronicsFactory Конкретная фабрика для создания продуктов в категории "Электроника"
type ElectronicsFactory struct{}

func (ef ElectronicsFactory) CreateProduct() Product {
	return &Phone{}
}

func (ef ElectronicsFactory) CreateShippingService() ShippingService {
	return &Courier{}
}

// ClothingFactory Конкретная фабрика для создания продуктов в категории "Одежда"
type ClothingFactory struct{}

func (cf ClothingFactory) CreateProduct() Product {
	return &Shirt{}
}

func (cf ClothingFactory) CreateShippingService() ShippingService {
	return &PostalService{}
}

// BooksFactory Конкретная фабрика для создания продуктов в категории "Книги"
type BooksFactory struct{}

func (bf BooksFactory) CreateProduct() Product {
	return &Book{}
}

func (bf BooksFactory) CreateShippingService() ShippingService {
	return &Bookstore{}
}

// Phone Реализация продукта - телефон
type Phone struct{}

func (p *Phone) GetInfo() string {
	return "Смартфон - продукт из категории 'Электроника'"
}

// Shirt Реализация продукта - рубашка
type Shirt struct{}

func (s *Shirt) GetInfo() string {
	return "Рубашка - продукт из категории 'Одежда'"
}

// Book Реализация продукта - книга
type Book struct{}

func (b *Book) GetInfo() string {
	return "Книга - продукт из категории 'Книги'"
}

// Courier Реализация службы доставки - курьерская служба
type Courier struct{}

func (c *Courier) Deliver() {
	fmt.Println("Товар доставляется курьерской службой.")
}

// PostalService Реализация службы доставки - почта
type PostalService struct{}

func (p *PostalService) Deliver() {
	fmt.Println("Товар доставляется почтовой службой.")
}

// Bookstore Реализация службы доставки - книжный магазин
type Bookstore struct{}

func (bs *Bookstore) Deliver() {
	fmt.Println("Товар доставляется книжным магазином.")
}

// В этом примере мы рассматриваем использование шаблона "Абстрактная фабрика" в интернет-магазине.
// Фабрики создают объекты продуктов разных категорий и связанные с ними службы доставки.
// У нас есть абстрактный интерфейс `ProductFactory`, который определяет методы `CreateProduct()`
// и `CreateShippingService()` для создания объектов продуктов и связанных служб доставки соответственно.
// Абстрактный интерфейс `Product` определяет метод `GetInfo()`, который возвращает информацию о продукте,
// а интерфейс `ShippingService` определяет метод `Deliver()` для доставки товаров.
// Мы создаем несколько конкретных фабрик, таких как `ElectronicsFactory`, `ClothingFactory` и `BooksFactory`,
// каждая из которых реализует методы фабрики для создания продуктов конкретных категорий и связанных служб доставки.
// Каждый конкретный продукт, такой как `Phone`, `Shirt` и `Book`, реализует метод `GetInfo()`,
// который возвращает информацию о конкретном продукте. Каждая конкретная служба доставки,
// такая как `Courier`, `PostalService` и `Bookstore`, реализует метод `Deliver()` для доставки товаров.
// В функции `main` мы создаем экземпляры конкретных фабрик и используем их для создания объектов продуктов
// и связанных с ними служб доставки. Затем мы вызываем методы `GetInfo()` для каждого продукта для получения информации о нем,
// а затем методы `Deliver()` для каждой службы доставки, чтобы показать, как они доставляют товары.
// Этот пример демонстрирует, как шаблон "Абстрактная фабрика" может быть использован в интернет-магазине
// для создания связанных объектов продуктов и соответствующих служб доставки в зависимости от категории товара.

func main() {
	// Создание фабрики для продуктов категории 'Электроника'
	electronicsFactory := ElectronicsFactory{}
	electronicsProduct := electronicsFactory.CreateProduct()
	electronicsShippingService := electronicsFactory.CreateShippingService()
	fmt.Println(electronicsProduct.GetInfo())
	electronicsShippingService.Deliver()

	// Создание фабрики для продуктов категории 'Одежда'
	clothingFactory := ClothingFactory{}
	clothingProduct := clothingFactory.CreateProduct()
	clothingShippingService := clothingFactory.CreateShippingService()
	fmt.Println(clothingProduct.GetInfo())
	clothingShippingService.Deliver()

	// Создание фабрики для продуктов категории 'Книги'
	booksFactory := BooksFactory{}
	booksProduct := booksFactory.CreateProduct()
	booksShippingService := booksFactory.CreateShippingService()
	fmt.Println(booksProduct.GetInfo())
	booksShippingService.Deliver()
}
