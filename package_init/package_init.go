package packageinit

import "fmt"

var connection string

func init() {
	/*
		- func init
		- func init akan selalu dijalankan pertama kali, ketika package digunakan
		- func ini akan dijalankan sebelum func main

		- dalam kasus ini func init akan memberikan value ke variable langsung
	*/
	connection = "MySQL"
	fmt.Println("func init tanpa memanggil func di main")
}

func GetDatabase() string {
	return connection
}
