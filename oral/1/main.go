package main

import (
	"fmt"
	"strings"
)

func main() {
	strSlice := []string{"Hi", " ", "Gopher", "!"}
	fmt.Println(strings.Join(strSlice, ""))
}
