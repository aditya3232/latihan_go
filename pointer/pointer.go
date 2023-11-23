package main

import "fmt"

/*
	- pointer
	- secara default semua variable itu dipassing by value, bukan by reference
*/
/*
	- passing by value
	- passing by value = variable baru akan menerima copyan dari variable asli, jadi perubahan pada data di variable baru tidak akan pengaruh ke variable asli

	- contoh :
	variableA := 10
	variableB := variableA
	variableB = 20
	fmt.Println(variableA) // 10
	fmt.Println(variableB) // 20
*/
/*
	- passing by reference
	- anda dapat menggunakan passing by reference dengan pointer *
	-  Saat Anda mengirimkan pointer ke suatu fungsi, Anda mengirimkan alamat memori dari variabel tersebut.
	- Dengan demikian, perubahan yang dilakukan pada nilai melalui dereferensi pointer akan mempengaruhi variabel aslinya.

	- Contoh :
	variableA := 10
	variableB := &variableA
	*variableB = 20
	fmt.Println(variableA) // 20
	fmt.Println(variableB) // 20

	- & (ampersand) Simbol ini digunakan untuk mendapatkan alamat memori dari suatu variabel. Dalam konteks kode Anda, &variableA mengembalikan alamat memori dari variabel variableA
	- * (asterisk) Simbol ini digunakan untuk mendapatkan nilai dari alamat memori yang ditunjuk oleh pointer. Dalam konteks kode Anda, *variableB mengembalikan nilai dari alamat memori yang ditunjuk oleh pointer variableB.

*/
/*
	- new
	- new adalah function yang berfungsi untuk membuat pointer baru
	- sama dengan &
	- tapi bedanya hanya mengembalikan alamat memorinya saja, datanya tidak diambil
	- Penting untuk dicatat bahwa penggunaan new umumnya tidak disarankan di Go, dan biasanya lebih baik menggunakan literal struct atau inisialisasi langsung.
*/

type Address struct {
	Kota     string
	Provinsi string
	Negara   string
}

func main() {

	address1 := Address{
		Kota:     "Malang",
		Provinsi: "Jawa_Timur",
		Negara:   "Indonesia",
	}

	address2 := address1
	address2.Kota = "Surabaya"

	fmt.Println(address2) // {Surabaya Jawa_Timur Indonesia}
	fmt.Println(address1) // data tetap sama, tidak berubah {Malang Jawa_Timur Indonesia}

	// pointer
	address3 := &address1 // penerapan pointer, bisa ditulis seperti ini => var address3 *Address = &address1

	/*
		- bisa tanpa asterik di kasus ini => (*address3).Kota = "Pasuruan"
		- jika kita print maka provinsi dan negara tetap ikut ya, karena ampersand mengambil alamat * datanya
		- kalau mau ambil alamat memorinya saja gunakan new
	*/
	// address3.Kota = "Pasuruan"

	*address3 = Address{
		Kota:     "Pasuruan",
		Provinsi: "Jawa_Timur",
		Negara:   "Indonesia",
	}

	fmt.Println(address3) // &{Pasuruan Jawa_Timur Indonesia}
	fmt.Println(address1) // data berubah {Pasuruan Jawa_Timur Indonesia}

	// new
	// tidak disarankan
	address4 := new(Address)
	*address4 = Address{
		Kota:     "Bandung",
		Provinsi: "Jawa_Barat",
		Negara:   "Indonesia",
	}

	fmt.Println(address4)
	fmt.Println(*address3)
}
