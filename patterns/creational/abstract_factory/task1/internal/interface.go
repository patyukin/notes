package internal

type NewsFactory interface {
	CreateNews() News
	CreateNewsSource() NewsSource
}

type News interface {
	GetTitle() string
	GetBody() string
	GetCategory() string
}

type NewsSource interface {
	GetName() string
	GetURL() string
}
