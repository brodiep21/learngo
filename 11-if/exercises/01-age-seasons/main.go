// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

type user struct {
	name string
	pass string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage : [Username] [Password]`)
		return
	}
	msg := os.Args[1:]
	u := user{
		name: msg[0],
		pass: msg[1],
	}

	if u.name != "Brodie" {
		fmt.Printf("Access denied for %s\n", u.name)
	}
	if u.pass != "pogwart" {
		fmt.Printf("invalid password %s\n", u.pass)
	} else if u.name == "Brodie" && u.pass == "pogwart" {
		fmt.Printf("Access granted for %s\n", u.name)
	}
}
