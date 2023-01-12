package main

import (
	"fmt"
)

type User struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Otch         string `json:"otch"`
	Registration string `json:"registration"`
}

type Us struct {
}

var users []User

func main() {

	fmt.Println(ret("Ali"))

}

func ret(s string) User {

	us1 := User{
		Family:       "Isaev",
		Name:         "Ali",
		Otch:         "Alievich",
		Registration: "Ss",
	}

	us2 := User{
		Family: "Petrov",
		Name:   "Petr",
		Otch:   "Petrovich",
	}

	users = append(users, us1)
	users = append(users, us2)

	d := users[1]
	fmt.Println(d)
	var sp int

	for i, user := range users {
		if user.Name == s {
			sp += i
			break
		}
	}
	fmt.Println(users[sp].Name)
	return User{}
}
