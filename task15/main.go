package main

import (
	"GolangBeginning/task15/geometry"
	"fmt"
)

func measure(g geometry.Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := geometry.Rect{Width: 3, Height: 4}
	c := geometry.Circle{Radius: 5}
	measure(r)
	measure(c)
}
