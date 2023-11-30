package main

import (
	"encoding/json"
	"fmt"
)

/*
	- berikan json tag agar key nya menjadi snake case,
	- dengan menambahkan json tag, maka ketika berkomunikasi dengan data json, key yang dipanggil adalah tagnya
*/

type Address struct {
	Region   string `json:"region"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type Customer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Hobbies   []string  `json:"hobbies"`
	Address   []Address `json:"address"`
}

func main() {
	/*
		- json marshal digunakan untuk encode data golang ke json,
		- data yg dimarshal akan menjadi bytes
		- maka dari itu harus di convert ke string
	*/

	address := []Address{
		{
			Region:   "Indonesia",
			City:     "Malang",
			Province: "Jawa Timur",
		},
		{
			Region:   "Indonesia",
			City:     "Surabaya",
			Province: "Jawa Timur",
		},
	}

	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Aditya",
		Age:       27,
		Hobbies:   []string{"game", "study", "watch movie"},
		Address:   address,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("error marshal", err)
	}

	fmt.Println("data asli", string(bytes))
}
