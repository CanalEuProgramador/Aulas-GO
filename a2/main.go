package main

import "fmt"

func main() {
	number := 20
	n := 10

	if n := 5; n > 0 {
		fmt.Println(n)
		fmt.Println("MAIOR QUE 0")
	} else {
		fmt.Println("MENOR QUE 0")
	}

	fmt.Println(n)

	switch number {
	case 2:
		fmt.Println("Number 2")
	case 5:
		fmt.Println("Number 5")
	default:
		fmt.Println("DEFAULT")
	}
}
