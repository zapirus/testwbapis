package main

import (
	"encoding/json"
	"fmt"
)

type Ud struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Otch         string `json:"otch"`
	Registration string `json:"registration"`
}

var uds []Ud

func main() {

	fmt.Println(rt("Petrov"))

}

func rt(s string) Ud {

	us1s := Ud{
		Family:       "Isaev",
		Name:         "Ali",
		Otch:         "Alievich",
		Registration: "Ss",
	}

	us2s := Ud{
		Family: "Petrov",
		Name:   "Petr",
		Otch:   "Petrovich",
	}

	uds = append(uds, us1s)
	uds = append(uds, us2s)

	mp := make(map[string]interface{})

	for _, user := range uds {
		if user.Name == s {
			mp["Name"] = s
		} else if user.Otch == s {
			mp["Otch"] = s
		} else if user.Family == s {
			mp["Family"] = s
		} else if user.Registration == s {
			mp["Registration"] = s
		}
	}
	jsonbody, err := json.Marshal(mp)
	if err != nil {
		// do error check
		fmt.Println(err)
	}
	student := Ud{}

	if err = json.Unmarshal(jsonbody, &student); err != nil {
		// do error check
		fmt.Println(err)
		return Ud{}
	}
	fmt.Printf("%#v\n", student)
	return student
}
