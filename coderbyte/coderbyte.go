package main

import "fmt"

func BracketCombinations(num int) int {
	if num == 0 {
		return 1
	}

	num = 0
	for i := 0; i < num; i++ {
		num += BracketCombinations(i) * BracketCombinations(num-i-1)
	}

	return num
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketCombinations(2))

}
