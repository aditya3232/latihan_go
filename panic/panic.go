package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
}

func jalankanBarang(grade string) (result string) {
	if grade == "Good" {
		result = "barang lanjut"
	} else {
		panic("barang tidak boleh lewat")
	}

	return result
}

func cekKualitasBarang(grade string, jalankanBarang func(string) string) (result string) {
	result = jalankanBarang(grade)

	return result

}

func RecoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	/*
		- panic fungsi untuk menghentikan program
		- program akan berhenti, tapi defer tetap dieksekusi
		- agar aplikasi tetap berjalan maka harus ada defer recover
	*/

	defer endApp()
	defer RecoverPanic()
	runApp := cekKualitasBarang("Sobek", jalankanBarang)
	fmt.Println(runApp)
}
