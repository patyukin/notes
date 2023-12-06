package main

import (
	"fmt"
)

var _ AbstractFactory = (*MacOSFactory)(nil)
var _ AbstractFactory = (*WindowsFactory)(nil)

// AbstractFactory интерфейс абстрактной фабрики
type AbstractFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Button интерфейс кнопки
type Button interface {
	Render() string
}

// Checkbox интерфейс флажка
type Checkbox interface {
	Render() string
}

// MacOSFactory конкретная фабрика для создания элементов интерфейса в стиле macOS
type MacOSFactory struct{}

func (f MacOSFactory) CreateButton() Button {
	return MacOSButton{}
}

func (f MacOSFactory) CreateCheckbox() Checkbox {
	return MacOSCheckbox{}
}

// WindowsFactory конкретная фабрика для создания элементов интерфейса в стиле Windows
type WindowsFactory struct{}

func (f WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (f WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// MacOSButton конкретная реализация кнопки для macOS
type MacOSButton struct{}

func (b MacOSButton) Render() string {
	return "Rendered a button in macOS style."
}

// MacOSCheckbox конкретная реализация флажка для macOS
type MacOSCheckbox struct{}

func (c MacOSCheckbox) Render() string {
	return "Rendered a checkbox in macOS style."
}

// WindowsButton конкретная реализация кнопки для Windows
type WindowsButton struct{}

func (b WindowsButton) Render() string {
	return "Rendered a button in Windows style."
}

// WindowsCheckbox конкретная реализация флажка для Windows
type WindowsCheckbox struct{}

func (c WindowsCheckbox) Render() string {
	return "Rendered a checkbox in Windows style."
}

func main() {
	// Создание фабрики для стиля macOS
	macOSFactory := MacOSFactory{}
	macOSButton := macOSFactory.CreateButton()
	macOSCheckbox := macOSFactory.CreateCheckbox()
	fmt.Println(macOSButton.Render())   // Output: Rendered a button in macOS style.
	fmt.Println(macOSCheckbox.Render()) // Output: Rendered a checkbox in macOS style.

	// Создание фабрики для стиля Windows
	windowsFactory := WindowsFactory{}
	windowsButton := windowsFactory.CreateButton()
	windowsCheckbox := windowsFactory.CreateCheckbox()
	fmt.Println(windowsButton.Render())   // Output: Rendered a button in Windows style.
	fmt.Println(windowsCheckbox.Render()) // Output: Rendered a checkbox in Windows style.
}
