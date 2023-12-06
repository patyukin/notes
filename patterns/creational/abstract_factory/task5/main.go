package main

type Notebook interface {
	Image() string
	Characteristics() string
	Price() string
}

type HPNotebook struct{}

func (h *HPNotebook) Image() string {
	return "HP image"
}

func (h *HPNotebook) Characteristics() string {
	return "HP characteristics"
}

func (h *HPNotebook) Price() string {
	return "HP price"
}

type DellNotebook struct{}

func (d *DellNotebook) Image() string {
	return "Dell image"
}

func (d *DellNotebook) Characteristics() string {
	return "Dell characteristics"
}

func (d *DellNotebook) Price() string {
	return "Dell price"
}

type AbstractFactory interface {
	CreateNotebook() Notebook
}

type HPFactory struct{}

func (h *HPFactory) CreateNotebook() Notebook {
	return &HPNotebook{}
}

type DELLFactory struct{}

func (d *DELLFactory) CreateNotebook() Notebook {
	return &DellNotebook{}
}

func main() {
	hpFactory := &HPFactory{}
	notebook := hpFactory.CreateNotebook()
	println(notebook.Image())
	println(notebook.Characteristics())
	println(notebook.Price())

	dellFactory := &DELLFactory{}
	notebook = dellFactory.CreateNotebook()
	println(notebook.Image())
	println(notebook.Characteristics())
	println(notebook.Price())
}
