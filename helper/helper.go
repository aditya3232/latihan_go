package helper

/*
	- package & import
	- package untuk run aplikasi wajib package main, dan func main(){}
	- secara standar, file golang hanya dapat mengakses file lainnya yang berada dalam package yang sama
	- untuk mengakses file lainnya yang berada di package yang berbeda, kita harus menggunakan import

*/

func SayHello(text string) (whatSay string) {
	whatSay = "hELLo " + text
	return whatSay
}
