package main

import (
	"fmt"
	"strconv"
)

var fizzBuzzMap = map[int]string{
	3: "Fizz",
	5: "Buzz",
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(getFizzBuzz(i))
	}
}

func getFizzBuzz(number int) string {
	output := ""

	for key, value := range fizzBuzzMap {
		if number%key == 0 {
			output += value
		}
	}

	if output == "" {
		output = strconv.Itoa(number)
	}

	return output
}
