package main

import "fmt"

func FizzBuzz(n int32) {
	for i := int32(1); i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
