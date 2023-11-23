package main

import "fmt"

type Alamat struct {
	Kota     string
	Provinsi string
	Negara   string
}

func ubahNegara(alamat *Alamat) {
	alamat.Negara = "Indonesia"

}

func main() {
	// data ini tidak akan merubah struct, karena tidak menggunakan pointer
	// kalau menggunakan pointer, maka akan merubah data struct
	inputAlamat := Alamat{
		Kota:     "Malang",
		Provinsi: "JawaTimur",
		Negara:   "",
	}

	ubahNegara(&inputAlamat)
	fmt.Println(inputAlamat)

}
