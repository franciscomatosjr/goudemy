package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.145
	var raio = 3.2
	// forma reduzida de criar um var
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	println(g, h, i)

}
