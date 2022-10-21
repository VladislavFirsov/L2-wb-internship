package main

import (
	"fmt"
	"math"
)

type figures interface {
	accept(visitor)
}

type visitor interface {
	visitForSquare(*square)
	visitForTriangle(*triangle)
}

type square struct {
	side int
}

func (s *square) accept(visitor visitor) {
	visitor.visitForSquare(s)
}

type triangle struct {
	a, b, c int
}

func (t *triangle) accept(visitor visitor) {
	visitor.visitForTriangle(t)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = s.side * s.side
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	sp := float64((t.a + t.b + t.c) / 2)
	a.area = int(math.Sqrt(sp * (sp - float64(t.a)) * (sp - float64(t.b)) * (sp - float64(t.c))))

}
func main() {
	calc := areaCalculator{}
	sq := square{side: 4}
	calc.visitForSquare(&sq)
	fmt.Println(calc.area)
	tr := triangle{
		a: 2,
		b: 3,
		c: 4,
	}
	calc.visitForTriangle(&tr)
	fmt.Println(calc.area)
}
