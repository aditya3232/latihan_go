package main

import "fmt"

/*
	- interface adalah abstrak dari method
	- atau definisi dari method
	- biasanya interface digunakan sebagai kontrak
*/

type Hello interface {
	HelloGuys() (say string)
}

type hello struct {
	name string
}

func (h hello) HelloGuys() (say string) {
	say = "Hi, nama saya adalah " + h.name

	return say
}

func main() {
	x := hello{
		name: "Ichsan",
	}

	fmt.Println(x.HelloGuys())
}
