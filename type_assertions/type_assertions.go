package main

import "fmt"

/*
	- type assertions
	- kemampuan merubah tipe data menjadi tipe data yang diinginkan
	- biasanya digunakan untuk memberikan tipe data pada  interface{}
	- tapi tetap harus sadar sesuai data yang diberikan, misal datanya adalah string kita berikan type assertions int akan panic
	- agar lebih aman gunakan switch case, dalam melakukan type assertions biar gk panic
*/

func Print(p interface{}) (print any) {
	print = p
	return print
}

func main() {
	input := 1
	result := Print(input)

	// yang di case, itu adalah type assertionsnya
	switch v := result.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Printf("int %d", v)
	default:
		fmt.Println("unknown")
	}

}
