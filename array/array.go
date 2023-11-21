package main

import "fmt"

func main() {
	/*
		- jika memiliki tipe data yg sama bisa disimpan di array
		- tentukan daya tampung array
		- di array tidak ada append, append adanya di slice
	*/

	var names [3]string

	names[0] = "Raisa"
	names[1] = "Andi"
	names[2] = "Budi"

	fmt.Println(names)
	fmt.Println(names[2]) // ambil data array

	names[2] = "Adit" // merubah data

	fmt.Println(names)

	// deklarasi array langsung isi datanya
	var hobbies = [3]string{
		"Sleeping",
		"Eating",
		"Playing",
	}

	fmt.Println(hobbies)

	// panjang array, bukan jml data
	fmt.Println(len(hobbies))
}
