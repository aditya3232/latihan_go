package main

import (
	"fmt"
	"strings"
)

/*
	- beberapa fungsi string untuk mengolah data string
	- len("halo guys") => cek panjang string
	- "halo guys"[1] => ambil byte string diposisi mana
	- string("halo guys"[1]) => konversi byte ke string
	- strings.Contains() => mendeteksi apakah string (parameter kedua) merupakan bagian dari string (parameter pertama), hasilnya bool
	- strings.HasPrefix() => mendeteksi apakah string (parameter kedua) adalah awalan string dari (parameter pertama), hasilnya bool
	- strings.HasSuffix() => mendeteksi apakah string (parameter kedua) adalah akhiran string dari (parameter pertama), hasilnya bool
	- strings.Count() => mendeteksi berapa jumlah string (parameter kedua), didalam string (parameter pertama), hasilnya int
	- strings.Index() => mendeteksi posisi string pada (parameter kedua), didalam string (parameter pertama), hasilnya int
	- strings.Replace() => untuk replace string, text = "...", find = "", replace = "", repeat = 1
	- strings.Split() => tentukan pemisah di (parameter kedua), lalu terapkan di (parameter pertama)
	- strings.Join() => digunakan untuk menggabungkan slice string (parameter pertama), dengan string di (parameter kedua)
	- strings.ToLower() => mengubah huruf string ke lowercase
	- strings.ToUpper() => mengubah huruf string ke Uppercase
*/

func main() {
	fmt.Println("halo guys")

	fmt.Println(len("halo guys"))

	fmt.Println("halo guys"[1])

	fmt.Println(string("halo guys"[1]))

	stringContains := strings.Contains("Budi sukan main Genshin", "Genshin") // apakah kata genshin ada di parameter pertama ? (true)
	fmt.Println(stringContains)

	stringHasPrefix := strings.HasPrefix("Budi suka main Genshin", "Bud") // apakah kata Bud adalah awalan dari parameter pertama ? (true)
	fmt.Println(stringHasPrefix)

	stringHasSuffix := strings.HasSuffix("Budi suka main Genshin", "shin") // apakah kata shin adalah akhiran dari parameter pertama ? (true)
	fmt.Println(stringHasSuffix)

	stringCount := strings.Count("Budi suka main Genshin", "i") // berapa jumlah kata i didalam parameter pertama ? (3)
	fmt.Println(stringCount)

	stringIndex := strings.Index("Budi suka main Genshin", "main") // posisi kata main didalam parameter pertama ? (10)
	fmt.Println(stringIndex)

	stringReplace := strings.Replace("Budi suka main Genshin", "Budi", "Ani", 1) // replace kata Budi dengan Ani, 1x saja
	fmt.Println(stringReplace)

	stringSplit := strings.Split("Budi suka main Genshin", " ") // pisahkan kata dengan spasi, hasilnya slice [Budi suka main Genshin]
	fmt.Println(stringSplit)

	stringJoin := strings.Join([]string{"Budi", "suka", "main", "Genshin"}, "-") // gabungkan slice string dengan -, hasilnya string Budi-suka-main-Genshin
	fmt.Println(stringJoin)

	stringToLower := strings.ToLower("Budi suka main Genshin") // ubah huruf menjadi lowercase
	fmt.Println(stringToLower)

	stringToUpper := strings.ToUpper("Budi suka main Genshin") // ubah huruf menjadi Uppercase
	fmt.Println(stringToUpper)

}
