package main

import "fmt"

func main() {
	var adm Admin
	var su simpleUser

	su.Message()

	adm.Login()

	fmt.Println(adm.online)

	adm.Logout()

	fmt.Println(adm.online)
}

type User struct {
	name   string
	online bool
}

func (u *User) Login() {
	u.online = true
}

func (u *User) Logout() {
	u.online = false
}

type Admin struct {
	User
	ranking int
}

type simpleUser User

func (su simpleUser) Message() {
	fmt.Println("Olá")
}
