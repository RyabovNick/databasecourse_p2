package main

import (
	"math"

	"github.com/RyabovNick/databasecourse_2/golang/intro/3_structs/creature"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	cr := creature.New()
	cr.Eat()
}
