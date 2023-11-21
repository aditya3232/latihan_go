package main

import "fmt"

func main() {
	/*
		- slice adalah potongan dari array
		- tipe data slice ada 3
		- pointer = petunjuk data pertama dari slice, index awal slice
		- length = panjang dari slice, total index slice
		- capacity = kapasitas dari slice, dimana length tidak boleh lebi dari capacity, total index awal sampai index akhir dari array
	*/

	/*
		- membuat slice dari array
		- array[low:high] = ambil potongan array dari index low ke index sebelum high
		- array[low:] = index low ke index akhir
		- array[high:] = index 0 sampai index sebelum high
		- array[:] = index 0 sampai index akhir
	*/

	/*
		- fungsi slice
		- len(slice) = panjang slice
		- cap(slice) = kapasitas slice
		- append(slice, data) = menambah data slice, jika kapsitas penuh, maka akan membuat array baru
		- make([]typeData, length, capacity) = membuat slice baru tanpa perlu membuat array, array otomatis terbuat
		- copy(destination, source) = menyalin slice dari source ke destination
	*/

	/*
		- digolang jarang pakai array, lebih sering slice
	*/

	array := [6]string{
		"Raisa",
		"Adit",
		"Budi",
		"Raka",
		"Rafi",
	}

	slice1 := array[1:4]
	/*
		- append slice, ketika melebihi kapasitas, dia akan tetap bisa, karena akan membuat array baru,
		- namun jika tidak meleibihi kapasitas dia tetap pakai array lama, dan ketika append, maka array lama juga akan berubah
	*/
	slice2 := append(slice1, "Rika", "Tomo", "Diqi")
	slice2[0] = "AditBaru" // merubah data slice

	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice2)

	newSlice := make([]string, 3, 4) // length harus lebih kecil dari capacity

	newSlice[0] = "Dota"
	newSlice[1] = "CS"
	newSlice[2] = "GenshinImpact"
	// newSlice[3] = "Valorant" // error, karena length lebih kecil dari capacity

	newSlice2 := append(newSlice, "Valorant")

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := make([]string, 7, 10) // slice belum ada datanya, bisa dibilang inisialisasi slice
	fromSlice[0] = "Senin"
	fromSlice[1] = "Selasa"
	fromSlice[2] = "Rabu"
	fromSlice[3] = "Kamis"
	fromSlice[4] = "Jumat"
	fromSlice[5] = "Sabtu"
	fromSlice[6] = "Minggu"

	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // length dan capacity destination harus sama seperti source

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// slice dengan data langsung
	newSlice3 := []uint8{
		90,
		80,
		70,
	}
	fmt.Println(newSlice3)

}
