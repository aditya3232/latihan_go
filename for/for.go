package main

import "fmt"

func main() {
	/*
		- for
		- pakai ; bukan ,
	*/

	/*
		- for range
		- digunakan untuk iterasi terhadap data collection array, slice, atau map
	*/

	/*
		- break = menghentikan perulangan
		- continue = menghentikan perulangan saat ini, dan melanjutkan ke perulangan berikutnya
	*/

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{"Raisa", "Adit", "Budi", "Raka", "Rafi"}

	for i := 0; i < len(slice); i++ { // kenapa kurang dari, karena length harus lebih kecil dari kapacity slice
		fmt.Println(slice[i])
	}

	for index, value := range slice { // tidak perlu pakai len
		fmt.Println(index, "=", value)
	}

	for _, value := range slice {
		if value == "Adit" {
			continue
		} else if value == "Raka" {
			break
		}
		fmt.Println(value)
	}

}
