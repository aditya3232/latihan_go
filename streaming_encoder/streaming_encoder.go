package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	- biasanya digunakan untuk menulis log
	- untuk membuat json di file, kita gunakan json.NewEncoder
*/

func main() {
	fileName := "log.log"

	// Open the file in append mode or create it if it doesn't exist
	writer, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error write", err)
	}

	defer writer.Close()

	encoder := json.NewEncoder(writer)

	data := map[string]interface{}{
		"level":     "info",
		"message":   "Success to get detail setting mcu",
		"timestamp": "18:50:00 30-11-2023",
	}

	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error encode", err)
	}

}
