package main

import "fmt"

func main() {
	/*
		- operasi perbandingan pasti akan menghasilkan boolean
		- >
		- <
		- >=
		- <=
		- ==
		- !=
	*/

	var (
		namasatu         = "budi"
		namadua          = "budi"
		namatiga         = "andini"
		angakasatu uint8 = 1
		angakadua  uint8 = 2

		perbandingan1 bool = namasatu == namadua
		perbandingan2      = namasatu != namatiga
		perbandingan3      = angakasatu > angakadua
	)

	fmt.Println(perbandingan1)
	fmt.Println(perbandingan2)
	fmt.Println(perbandingan3)
}
