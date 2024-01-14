package main

import "fmt"

type USPlug interface {
	USPlugIn()
}

type Socket struct {
}

func (s *Socket) Plug(plug USPlug) {
	fmt.Println("Пробуем что-то вставить в американскую розетку.")
	plug.USPlugIn()
}

type Plug struct {
}

func (p *Plug) USPlugIn() {
	fmt.Println("Американская вилка вставлена.")
}

// EuroPlug - не поддерживает интерфейс USPLug
type EuroPlug struct {
}

func (e *EuroPlug) EuroPlugIn() {
	fmt.Println("Евровилка вставлена.")
}

// Adapter  реализующий интерфейс USPlug
type Adapter struct {
	Euro *EuroPlug // Делаем встраивание, чтоб потом вызвать метод EuroPlugIn()
}

// USPlugIn В этом методе мы вызываем метод EuroPlugIn(),
func (a *Adapter) USPlugIn() {
	fmt.Println("Адаптер в американской розетке.")
	a.Euro.EuroPlugIn() // вызываем метод EuroPlugIn(), тут как бы мы делаем некий переходник(адаптер),
	// чтобы соответствовать интерфейсу
}

func main() {
	socket := &Socket{}
	socket.Plug(&Plug{})

	euroPlug := &EuroPlug{}
	adapter := &Adapter{Euro: euroPlug}

	socket.Plug(adapter)
}
