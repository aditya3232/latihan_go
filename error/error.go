package main

import (
	"errors"
	"fmt"
)

/*
	- untuk membuat error, go sudah ada packagenya, jadi tidak perlu bikin manual
	- tapi jika ingin membuat custom error juga bisa
*/

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}

}

func main() {
	nilai := 20
	pembagi := 0

	hasil, err := Pembagian(nilai, pembagi)
	if err != nil {
		fmt.Println("error di fungsi pembagian : ", err.Error())
	}

	fmt.Println(hasil)
}
