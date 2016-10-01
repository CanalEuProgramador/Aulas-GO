package main

import "fmt"

var i = 10

func main() {
	fmt.Println(soma(12, 13, 30, 100))
	fakeLoop()
}

func soma(numbers ...int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total
}

func fakeLoop() {
	fmt.Println(i)
	i--
	if i > 0 {
		fakeLoop()
	}
}
