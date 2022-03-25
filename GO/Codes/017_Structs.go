package main

import (
	"fmt"
)

type Point struct {
	x float32
	y float32
}

type Triangle struct {
	a Point
	b Point
	c Point
}

type Circle struct {
	radius float32
	*Point // you can name it center if you want
}

func main() {
	pt := Point{4, 5}
	fmt.Println("Point")
	fmt.Println(pt)
	fmt.Println(pt.x)
	fmt.Println()

	tri := Triangle{b: pt, c: Point{6, 8}}
	tri.a = Point{3, 2}
	fmt.Println("Triangle")
	fmt.Println(tri)
	fmt.Println(tri.a)
	fmt.Println(tri.b.x)
	fmt.Println()

	cir := Circle{6, &Point{7, 1}}
	fmt.Println("Circle")
	fmt.Println(cir)
	fmt.Println(cir.radius)
	fmt.Println(cir.x)
}
