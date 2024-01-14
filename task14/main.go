package main

import (
	"fmt"
)

func main() {
	var a any = 1
	var b any = "str"
	var c any = true
	var d any = make(chan struct{})
	check(a)
	check(b)
	check(c)
	check(d)
}

// Создаем функцию, где через .(type) определяем тип переменой в switch
func check(v any) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan")
	}
}

// Здесь с помощью пакета reflect определяем тип переменной
//func check(v any) {
//	switch reflect.ValueOf(v).Kind() {
//	case reflect.Int:
//		fmt.Println("int")
//	case reflect.String:
//		fmt.Println("string")
//	case reflect.Bool:
//		fmt.Println("bool")
//	case reflect.Chan:
//		fmt.Println("chan")
//	}
//}
