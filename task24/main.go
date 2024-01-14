package main

import (
	"fmt"
	"math"
)

// Pointer точка с координатами x y
type Pointer struct {
	x float64
	y float64
}

// Создаем точку
func createPointers(x, y float64) Pointer {
	return Pointer{
		x: x,
		y: y,
	}
}

// Высчитываем дистанцию между точками
func (p Pointer) distance(p2 Pointer) float64 {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := createPointers(2.0, 3.0)
	p2 := createPointers(4.0, 1.0)

	fmt.Printf("%.2f", p1.distance(p2))
}
