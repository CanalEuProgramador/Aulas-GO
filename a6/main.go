package main

import "fmt"

func main() {
	sl := make([]int, 2, 3)

	sl[0] = 90
	sl[1] = 80

	sl2 := sl[:]

	sl = append(sl, 1, 60)

	sl2[0] = 200

	fmt.Println(sl)
	fmt.Println(sl2)
	fmt.Println(len(sl))
}
