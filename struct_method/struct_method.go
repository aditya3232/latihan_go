package main

import "fmt"

/*
	- struct adalah tipe data, sehingga bisa digunakan sebagai parameter untuk fungsi
	- fungsi yang memiliki struct, namanya adalah method
*/

type Customer struct {
	Nama   string
	Alamat string
	Umur   int32
}

// ini adalah method
func (customer Customer) KatakanHai(namaPenjual string) {
	fmt.Println("Hai, Saya", namaPenjual, "selamat datang", customer.Nama, customer.Alamat)
}

// ini adalah function
func main() {
	// Nama => menampung data struct
	customer := Customer{
		Nama:   "Ichsan Ashiddiqi",
		Alamat: "Malang",
	}

	namaPenjual := "Fuad"

	// memanggil method
	customer.KatakanHai(namaPenjual)

}
