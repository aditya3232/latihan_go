package main

import (
	"fmt"
	"latihan-go/helper"

	// packageinit "latihan-go/package_init"
	_ "latihan-go/package_init"
)

func main() {
	say := "Adit"
	result := helper.SayHello(say)
	fmt.Println(result)

	/*
		- ketika ingin menjalankan func init() di package, tanpa harus memanggil salah satu func di main,
		- secara default golang akan komplain, karena import tanpa menggunakan func nya,
		- maka dari itu digunakan blank identifier di import packagenya, ketika hanya ingin run func init, tanpa menggunakannya di main
	*/
	// fmt.Println(packageinit.GetDatabase())
}
