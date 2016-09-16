package main

import "fmt"

func main() {
	result, msg := soma(10, 20)
	if msg != "" {
		fmt.Println(msg)
	}

	fmt.Println("O resultado da soma é:", result)
}

func soma(n1, n2 int) (int, string) {
	fmt.Println("Iniciando soma")

	result := n1 + n2
	if result > 10 {
		return result, "Esse valor é maior que 10"
	}

	return result, ""
}
