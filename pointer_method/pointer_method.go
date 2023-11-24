package main

import "fmt"

/*
	- sangat direkomendasikan menggunakan pointer method
	- sehingga tidak boros memory karena selalu diduplikasi ketika memanggil method
	- untuk membuat pointer method cukup memberikan tanda asterik di struct nya
*/

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr ." + man.Name // nilai asli ada Mr . nya

}

func main() {
	johnWick := Man{
		Name: "John Wick",
	}

	johnWick.Married()

	/*
		- jika tidak menggunakan pointer
		- data struct adalah John Wick; ini adalah copyan struct asli
		- struct asli tidak berubah; tetap Mr .

		- kalau menggunakan pointer datanya jadi Mr .John Wick
	*/
	fmt.Println(johnWick.Name)

}
