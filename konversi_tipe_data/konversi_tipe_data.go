package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai8 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi byte dari ambil string
	// byte itu uint8
	fmt.Println("ichsan"[0])
	fmt.Println(string("ichsan"[0]))
}
