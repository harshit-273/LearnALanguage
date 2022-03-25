package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (c *Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float32 {
	return r.length * r.breadth
}

func main() {
	c1 := Circle{10}
	r1 := Rectangle{4, 5}
	shapes := []Shape{&c1, &r1}

	for _, sh := range shapes {
		fmt.Println(sh.area())
	}
}
