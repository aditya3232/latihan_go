package main

import "fmt"

/*
	- di golang ketika sebuah variable atau parameter dengan tipe tertentu, maka secara otomatis ada data defaultnya
	- di golang, nil atau kosong itu hanya dapat digunakan dibeberapa tipe data yaitu :

	- interface
	- function
	- map
	- slice
	- pointer
	- channel
*/

func NewMap(name string) (result map[string]string) {
	if name == "" {
		return nil // ini contoh penggunaan nil di map
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	myNameIs := "aditya"
	print := NewMap(myNameIs)

	fmt.Println(print["name"]) // pemanggilan key dari map
}
