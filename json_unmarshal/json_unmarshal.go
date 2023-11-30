package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Region   string
	City     string
	Province string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
	Address   []Address
}

func main() {
	/*
		- json unmarshal digunakan untuk decode data json ke golang,
		- data yg diunmarshal harus di convert ke bytes, jsonString, jsobBytes
		- jangan lupa gunakan pointer agar data struct berubah
		- parameter nya json.Unmarshal(jsonBytes, struct)
	*/

	jsonString := `{
	"FirstName": "Muhammad",
	"LastName": "Aditya",
	"Age": 27,
	"Hobbies": [
		"game",
		"study",
		"watch movie"
	],
	"Address": [
		{
		"Region": "Indonesia",
		"City": "Malang",
		"Province": "Jawa Timur"
		},
		{
		"Region": "Indonesia",
		"City": "Surabaya",
		"Province": "Jawa Timur"
		}
	]
	}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		fmt.Println("error unmarshal customer", err)
	}

	fmt.Println("data asli", customer)
	fmt.Println("data asli", customer.FirstName)
	fmt.Println("data asli", customer.Hobbies)
	fmt.Println("data asli", customer.Address)

	/*
		- unmarshall address saja
		- disini pakai slice, karena ada 2 object di Address => address := &[]Address{}
	*/
	jsonAddressString := `
	[
		{
		"Region": "Indonesia",
		"City": "Malang",
		"Province": "Jawa Timur"
		},
		{
		"Region": "Indonesia",
		"City": "Surabaya",
		"Province": "Jawa Timur"
		}
	]
	`
	jsonBytesAddress := []byte(jsonAddressString)

	address := &[]Address{}

	err = json.Unmarshal(jsonBytesAddress, address)
	if err != nil {
		fmt.Println("err unmarshall json address", err)
	}

	fmt.Println("data address", address)

}
