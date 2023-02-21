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
type users struct {
	users []user
}

var u = users{
	[]user{
		{
			name: "jack",
			pass: "butter",
		},
		{
			name: "Brodie",
			pass: "buffalo",
		},
	},
}

const (
	accDenied = "Access denied for %s\n"
	invPass   = "invalid password %s\n"
	accessOK  = "Access granted for %s\n"
)

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
	checkUser(u)

}

func checkUser(msg user) {
	if msg.name != u.users[0].name && msg.name != u.users[1].name {
		fmt.Printf(accDenied, msg.name)
		return
	} else if msg.pass != u.users[0].pass && msg.pass != u.users[1].pass {
		fmt.Printf(invPass, msg.pass)
		return
	} else if msg.name == u.users[0].name && msg.pass == u.users[0].pass || msg.name == u.users[1].name && msg.pass == u.users[1].pass {
		fmt.Printf(accessOK, msg.name)
		return
	}

}
