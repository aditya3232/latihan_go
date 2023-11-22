package main

import "fmt"

func tesDefer() {
	fmt.Println("aplikasi selesai")
}

func main() {
	/*
		- defer adalah fungsi yang berjalan setelah sebuah fungsi selesai (berjalan diakhir)
		- defer pasti akan dieksekusi walaupun ada error di fungsi yg dieksekusi
	*/

	defer tesDefer() // dieksekusi terakhir

	aplikasiJalan := func() string {
		teks := "aplikasi jalan"
		return teks
	}

	tes := aplikasiJalan
	fmt.Println(tes())

}
