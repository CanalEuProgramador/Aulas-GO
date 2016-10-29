package main

import (
	"fmt"
)

func main() {
	u := User{
		"Eliseu",
		"zeucxb",
		true,
	}

	adm := Admin{
		User{
			"Roberto",
			"ro<3",
			true,
		},
		19,
	}

	ShowUserInfo(u)
	ShowUserInfo(adm)
}

func ShowUserInfo(u UsersInterface) {
	fmt.Println(u.Show())
}

type UsersInterface interface {
	Show() string
}

type User struct {
	name     string
	username string
	online   bool
}

func (u User) Show() string {
	return fmt.Sprintf("Hello, my name is %s, and my username is %s.", u.name, u.username)
}

type Admin struct {
	User
	age int
}

func (u Admin) Show() string {
	return fmt.Sprintf("Hello, my name is %s, and my username is %s. I`m a Adm!!!", u.name, u.username)
}
