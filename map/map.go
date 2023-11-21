package main

import "fmt"

func main() {
	/*
		- map
		- adalah kumpulan key-value, dimana key harus unik
		- jml data yg dimasukkan boleh sebanyak2 nya
	*/

	/*
		- fungsi map
		- len(map) = jml data
		- map[key] = ambil data
		- map[key] = value => mengubah data di map
		- make(map[typeKey]typeValue) = membuat map baru
		- delete(map, key) = menghapus data di map
	*/

	person := map[string]string{
		"nama": "adit",
		"hobi": "main_game",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])

	buahApel := make(map[string]string)

	buahApel["nama"] = "apel"
	buahApel["warna"] = "merah"

	fmt.Println(buahApel)

	buahApel["warna"] = "hijau"

	delete(buahApel, "warna")

	fmt.Println(buahApel)

}
