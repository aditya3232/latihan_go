package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		= kadang File JSON yang akan di decode itu berasal dari io.Reader, misal : file, network, request body
		- maka dari itu cara decodenya menggunakan NewDecoder
		- kemudian simpan di struct seperti biasa
	*/
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
