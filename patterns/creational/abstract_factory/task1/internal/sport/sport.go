package sport

import "task1/internal"

var _ internal.News = (*News)(nil)

type News struct{}

func (sn *News) GetTitle() string {
	return "Спорт"
}

func (sn *News) GetBody() string {
	return "Новости спорта"
}

func (sn *News) GetCategory() string {
	return "Категория: Спорт"
}

type Factory struct{}

func (sf *Factory) CreateNews() internal.News {
	return &News{}
}
