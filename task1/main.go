package main

import "fmt"

type Action struct { // Дочерняя структура
	hobby string
	Human // Встроенная родительская функция, теперь доступны поля
}

type Human struct { // Родительская структура
	Name    string
	Surname string
	Age     int
	Country string
}

func (h *Human) info() {
	fmt.Printf("I'm %s %s. I'm %d. I'm from %s \n", h.Name, h.Surname, h.Age, h.Country)
}

func (a *Action) activity() {
	fmt.Printf("I go to the %s\n", a.hobby)
}

func main() {
	human := Human{
		Name:    "Nikita",
		Surname: "Baryshnikov",
		Age:     22,
		Country: "Russia",
	}
	action := Action{
		hobby: "gym",
		Human: human,
	}
	//Метод Human, можем вызывать из структуры Action(т.к у структуры Action нет метода info)
	// Если бы у Action был такой метод, то был бы shadowing (тогда придется вызывать так: action.Human.info())
	action.info()
	action.activity() // Метод Action
}
