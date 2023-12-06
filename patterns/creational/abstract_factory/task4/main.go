package main

import "fmt"

type Button interface {
	Render() string
}

type QuantityInput interface {
	Render() string
}

type AddressInput interface {
	Render() string
}

type MarketAbstractFactory interface {
	CreateButton() Button
	CreateQuantityInput() QuantityInput
	CreateAddressInput() AddressInput
}

type RussiaButton struct{}

func (b *RussiaButton) Render() string {
	return "<button>Купить</button>"
}

type RussiaQuantityInput struct{}

func (q *RussiaQuantityInput) Render() string {
	return "<input type='number'>"
}

type RussiaAddressInput struct{}

func (a *RussiaAddressInput) Render() string {
	return "<input type='text'>"
}

type ChinaButton struct{}

func (b *ChinaButton) Render() string {
	return "<button>买</button>"
}

type ChinaQuantityInput struct{}

func (q *ChinaQuantityInput) Render() string {
	return "<input type='number'>"
}

type ChinaAddressInput struct{}

func (a *ChinaAddressInput) Render() string {
	return "<input type='text'>"
}

type RussiaFactory struct{}

func (f *RussiaFactory) CreateButton() Button {
	return &RussiaButton{}
}

func (f *RussiaFactory) CreateQuantityInput() QuantityInput {
	return &RussiaQuantityInput{}
}

func (f *RussiaFactory) CreateAddressInput() AddressInput {
	return &RussiaAddressInput{}
}

type ChinaFactory struct{}

func (f *ChinaFactory) CreateButton() Button {
	return &ChinaButton{}
}

func (f *ChinaFactory) CreateQuantityInput() QuantityInput {
	return &ChinaQuantityInput{}
}

func (f *ChinaFactory) CreateAddressInput() AddressInput {
	return &ChinaAddressInput{}
}

func main() {
	var factory MarketAbstractFactory

	button := factory.CreateButton()
	quantityInput := factory.CreateQuantityInput()
	addressInput := factory.CreateAddressInput()

	fmt.Println(button.Render())
	fmt.Println(quantityInput.Render())
	fmt.Println(addressInput.Render())
}
