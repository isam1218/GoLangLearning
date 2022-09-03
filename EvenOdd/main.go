package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		t := strconv.Itoa(number)

		if number%2 == 0 {
			fmt.Println(t + " is even")
		} else {
			fmt.Println(t + " is odd")
		}
	}
}
