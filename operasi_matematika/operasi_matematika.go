package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		i = 30
		x = 40

		c = a + b
		d = a - b
		e = a * b
		f = a / b
		g = a % b
	)

	// augmented assignments, artinya i = i + 10 + 5
	i += 10
	i += 5

	// unary operator
	x++ // 40 + 1
	x-- // 40 - 1

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(i)
	fmt.Println(x)
}
