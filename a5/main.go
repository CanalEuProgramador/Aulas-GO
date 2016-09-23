package main

import "fmt"

func main() {
	tabuada := [10]int{0, 5, 10}
	user := map[string]string{
		"name": "Eliseu",
		"nick": "Zeu Cxb",
	}
	slice := []int{0, 5, 10}

	slice = append(slice, 90, 110, 50)

	fmt.Println(tabuada)
	fmt.Println(user)
	fmt.Println(slice)
}
