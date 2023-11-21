package main

import "fmt"

func main() {
	/*
		- switch
		- biasanya digunakan untuk pengecekan kondisi dalam satu var
		- atau untuk operasi perbandingan
	*/

	keju := "11 hari"

	switch keju {
	case "2 hari":
		fmt.Println("sudah matang")
	case "6 hari":
		fmt.Println("sudah bagus")
	case "10 hari":
		fmt.Println("siap dijual")
	default:
		fmt.Println("keju busuk")
	}
}
