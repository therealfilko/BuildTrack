package main

import (
	"fmt"
	"strconv"
	"sync"
)

var fizzBuzzMap = map[int]string{
	3: "Fizz",
	5: "Buzz",
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

func fizzBuzzWorker(start int, end int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		ch <- getFizzBuzz(i)
	}
}

func main() {
	const numWorkers = 4
	const maxNumber = 100
	ch := make(chan string, maxNumber)
	var wg sync.WaitGroup

	rangeSize := maxNumber / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		if i == numWorkers-1 {
			end = maxNumber
		}
		wg.Add(1)
		go fizzBuzzWorker(start, end, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
