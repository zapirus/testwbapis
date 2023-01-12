package main

import "fmt"

type Used struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Otch         string `json:"otch"`
	Registration string `json:"registration"`
}

var use []Used

func main() {

	fmt.Println(rets("Ali"))

}

func rets(s string) Used {

	us1 := Used{
		Family:       "Isaev",
		Name:         "Ali",
		Otch:         "Alievich",
		Registration: "Ss",
	}

	us2 := Used{
		Family: "Petrov",
		Name:   "Petr",
		Otch:   "Petrovich",
	}

	use = append(use, us1)
	use = append(use, us2)

	for i, user := range use {
		if user.Name == s {
			return use[i]
		}
	}
	return Used{}
}
