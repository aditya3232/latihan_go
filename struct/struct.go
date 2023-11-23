package main

import "fmt"

/*
	- struct
	- adalah representasi data dalam aplikasi => atau bisa dibilang template data
	- misalnya dto atau data transfer object
	- nama struct dan fieldnya pakai Pascalcase, kapital didepan
	- jika struct kosong, maka akan diisi dengan default value dari tipe datanya
*/

type Customer struct {
	Nama   string
	Alamat string
	Umur   int32
}

func main() {
	var user Customer

	user.Nama = "Muhammad userya"
	user.Alamat = "Malang Indonesia"
	user.Umur = 27

	fmt.Println(user) // hasilnya adalah data yang dibungkus ke objek
	fmt.Println(user.Nama)
	fmt.Println(user.Alamat)
	fmt.Println(user.Umur)

	// cara kedua struct literals
	newUser1 := Customer{
		Nama:   "Ichsan Ashiddiqi",
		Alamat: "Malang Indonesia",
		Umur:   24,
	}

	fmt.Println(newUser1)
}
