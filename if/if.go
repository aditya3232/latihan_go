package main

import "fmt"

func main() {
	/*
		- if
		- yang dibandingkan adalah operasi perbandingan, atau operasi boolean
		- if lebih baik buat operasi boolean
		- jika hanya operasi perbandingan pake switch
	*/

	name := "budi"

	if name == "eko" {
		fmt.Println("halo eko")
	} else {
		fmt.Println("bukan eko")
	}

	keju := "6 hari"
	warna := "hijau"

	if keju == "2 hari" && warna == "kuning" {
		fmt.Println("keju matang")
	} else if keju == "6 hari" && warna == "kuning" {
		fmt.Println("keju sudah bagus")
	} else if keju == "10 hari" && warna == "kuning" {
		fmt.Println("siap dijual")
	} else {
		fmt.Println("keju busuk")
	}

}
