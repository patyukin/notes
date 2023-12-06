package technology

import "task1/internal"

type News struct{}

func (tn *News) GetTitle() string {
	return "Технологии"
}

func (tn *News) GetBody() string {
	return "Новости технологии"
}

func (tn *News) GetCategory() string {
	return "Категория: Технологии"
}

type NewsSource struct{}

func (tns *NewsSource) GetName() string {
	return "Технологии"
}

func (tns *NewsSource) GetURL() string {
	return "https://www.technology.com"
}

type Factory struct{}

func (tf *Factory) CreateNews() internal.News {
	return &News{}
}
func (tf *Factory) CreateNewsSource() internal.NewsSource {
	return &NewsSource{}
}
