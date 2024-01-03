package main

import (
	"fmt"
)

type in chan int32

var serverChan in = make(chan int32)

func oddEven(n int32) bool {
	return n%2 != 0
}

func server() {
	var oddChan, evenChan chan int32
	oddChan = make(chan int32)
	evenChan = make(chan int32)

	for {
		num := <-serverChan
		if oddEven(num) {
			oddChan <- num
		} else {
			evenChan <- num
		}
	}
}

func main() {
	go server()

	var oddChan, evenChan chan int32

	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range arr {
		serverChan <- num
	}

	fmt.Println("Odd numbers:")
	for num := range oddChan {
		fmt.Println(num)
	}

	fmt.Println("Even numbers:")
	for num := range evenChan {
		fmt.Println(num)
	}
}
