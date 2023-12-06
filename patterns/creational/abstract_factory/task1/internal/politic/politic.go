package politic

import "task1/internal"

var _ internal.News = (*PoliticsNews)(nil)

type PoliticsNews struct{}

func (pn *PoliticsNews) GetTitle() string {
	return "Политика"
}

func (pn *PoliticsNews) GetBody() string {
	return "Новости политики"
}

func (pn *PoliticsNews) GetCategory() string {
	return "Категория: Политика"
}

type PoliticsFactory struct{}

func (pf *PoliticsFactory) CreateNews() internal.News {
	return &PoliticsNews{}
}
