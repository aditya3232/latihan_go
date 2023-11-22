package main

import "fmt"

/*
	- adalah blok kode yang dibuat agar dapat digunakan berulang2
	- func dapat dipanggil di func lain
	- variadic func, adalah fungsi yang menyimpan parameter berikut : (angka ...int64)
	- func dapat disimpan didalam variable
	- maka dari itu filter juga bisa dijadikan parameter
	- ada juga anonymous function, atau function tanpa nama, biasanya function didalam variable atau parameter
	- ada juga recursive function, yaitu function yang memanggil function diri sendiri, tapi kayanya jarang dipake
*/

/*
	- closure, adalah kemampuan func untuk berinteraksi dengan data-data disekitarnya dalam scope yg sama
*/

func sayHello(parameter string) {
	fmt.Println(parameter) // parameter wajib diisi
}

func penjumlahan(a, b int64) int64 { // return itu wajib jika sudah di deklarasikan
	c := a + b

	return c
}

func rataRataNilai(a, b, c int64) (result, grade string) {
	penjumlahan := a + b + c
	rataRata := penjumlahan / 3

	/*
		- Sprintf = digunakan untuk print string, dengan berbagai parameter
		- parameter tersebut akan menjadi string juga
		- %s = string
		- %d = int
		- %f = float
		- %b = biner
		- %t = boolean
		- %c = unicode
		- %% = mencetak karakter persen
	*/
	result = fmt.Sprintf("rata-rata nilai anda adalah : %d", rataRata)

	if rataRata >= 80 {
		grade = "dan grade anda adalah : A"
	} else if rataRata >= 60 && rataRata <= 79 {
		grade = "dan grade anda adalah : B"
	} else if rataRata >= 50 && rataRata <= 59 {
		grade = "dan grade anda adalah : C"
	} else {
		grade = "dan grade anda adalah : D"
	}

	// default-nya || else-nya
	return result, grade

}

/*
	- variadic function
	- parameter diposisi terakhir, yg dapat menerima lebih dari satu input (varargs)
	- sebenernya bisa pakai slice, tapi lebih enak pakai variadic function, jika ada diposisi terakhir
*/
func multipleAll(parameter ...int64) (total int64) {
	total = 1

	for _, value := range parameter {
		/*
			- total = total * value
			- atau dalam kasus ini adalah => total=1×2×3×4×5=120
			- karena yang di for hanya valuenya saja
		*/
		total *= value
	}

	return total
}

/*
	- function value
	- function itu bisa disimpan di variable
	- berarti function bisa juga menjadi parameter
*/

func goodBye(input string) (say string) {
	say = input

	return say
}

type Filter func(string) string // type declarations untuk function

// function sebagai parameter
func filterSpam(kata string, filter Filter) (filteredSpam string) {
	filteredSpam = filter(kata)

	return filteredSpam
}

// ini adalah parameternya
func filter(kata string) (filteredSpam string) {
	if kata == "anjing" {
		filteredSpam = "****"
	} else {
		filteredSpam = kata
	}

	return filteredSpam
}

func main() {
	sayHello("Halo guys")
	sayHello("Apa kabar kalian semua")
	sayHello("")

	result := penjumlahan(40, 10)
	fmt.Println(result)

	rataRata, grade := rataRataNilai(80, 40, 0)
	fmt.Println(rataRata, grade)

	// jika tidak ingin menggunakan salah satu return
	// _, grade := rataRataNilai(80, 80, 80)
	// fmt.Println(grade)

	result2 := multipleAll(2, 3, 4, 5)
	fmt.Println(result2)

	// slice sebagai data variadic func jg bisa
	slice := []int64{6, 7, 8, 9}
	result3 := multipleAll(slice...)
	fmt.Println(result3)

	// function value
	iSay := "good bye my friend"
	whatUSay := goodBye         // jgn pakai parameter
	fmt.Println(whatUSay(iSay)) // disini baru pakai parameter

	/*
		- function sebagai parameter
		- disini func filter tidak perlu pakai return, karena returnnya digunakan di filterSpam
	*/
	filteredSpam := filterSpam("kodok", filter)
	fmt.Println(filteredSpam)

	// ini adalah anonymous function
	blacklist := func(kata string) (filter string) {

		if kata == "anjing" {
			filter = "***"
		} else {
			filter = kata
		}

		return filter
	}

	filterWithAnonymousFunction := filterSpam("anjing", blacklist)
	fmt.Println(filterWithAnonymousFunction)
}
