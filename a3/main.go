package main

import "fmt"

func main() {
	number := 20

	for number >= 0 {
		fmt.Println(number)
	}

	// cria um novo number local, diferente de 20, igual a 0
	for number := 0; number < 10; number++ {
		fmt.Println(number)
	}
}
