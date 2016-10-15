package main

import "fmt"

func main() {
	u := user{
		"Eliseu",
		19,
		true,
	}

	fmt.Println(u.Live)

	u.Morrer()

	fmt.Println(u.Live)
}

type user struct {
	Name string
	Age  int
	Live bool
}

func (u *user) Morrer() {
	fmt.Println("Faliceu")

	u.Live = false
}
