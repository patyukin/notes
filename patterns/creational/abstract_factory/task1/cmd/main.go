package main

import (
	"task1/internal/technology"
)

// Задача:
// Создать систему для генерации различных видов новостей на основе шаблона "Абстрактная фабрика".
// Необходимо разработать фабрики для генерации новостей разных категорий: "Политика", "Спорт" и "Технологии".
// Каждая фабрика должна иметь методы для создания объекта новости (`CreateNews`) и объекта источника новости (`CreateNewsSource`).
// Новости и источники новостей должны быть связаны между собой.

func main() {
	technologyFactory := technology.Factory{}

	technologyNews := technologyFactory.CreateNews()
	println(technologyNews.GetCategory())
	println(technologyNews.GetTitle())
	println(technologyNews.GetBody())

	technologyNewsSource := technologyFactory.CreateNewsSource()
	println(technologyNewsSource.GetName())
	println(technologyNewsSource.GetURL())
}
