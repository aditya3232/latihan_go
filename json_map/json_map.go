package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		- karena pakai map, maka tidak membutuhkan struct untuk menampung datanya
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

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		fmt.Println("error unmarshal map", err)
	}

	fmt.Println("data asli", result)
	fmt.Println("data asli", result["FirstName"])
	fmt.Println("data asli", result["Hobbies"])
	fmt.Println("data asli", result["Address"])

	/*
		- json marshal
		- disini pakai func dengan return bytes
	*/
	bytes := func() []byte {
		product := map[string]interface{}{
			"id":    "P0001",
			"name":  "MSI GF63",
			"Price": "20000000",
		}

		bytes, err := json.Marshal(product)
		if err != nil {
			fmt.Println("error marshal map", err)
			return nil
		}

		return bytes
	}()

	// print return anonymous function yg disimpan di variable
	fmt.Println(string(bytes))

}
