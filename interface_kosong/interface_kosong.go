package main

import "fmt"

/*
	- interface kosong adalah sebuah deklarasi yg tidak memiliki deklarasi method satupun,
	- sehingga semua tipe data dapat menjadi implementasinya
	- interface{}, juga memiliki type alias any
*/

func Print(p interface{}) (print any) {
	print = p
	return print
}

func main() {
	input := true
	cetakInput := Print(input)

	fmt.Println(cetakInput)
}
