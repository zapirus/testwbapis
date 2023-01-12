package main

import (
	"fmt"
	"strings"
)

func main() {
	sg("/posts/sss")

}

func sg(s string) {
	sd := strings.Split(s, "/")
	fmt.Println(sd[2])
	for i, s2 := range sd {
		fmt.Println(i, s2)
	}

}
