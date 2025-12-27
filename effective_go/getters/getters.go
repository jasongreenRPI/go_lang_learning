package main

import (
	"fmt"
)

type User struct {
	id int
	name string
}

func (u User) Name() string {
	return u.name
}

func main() {
	user := User{
		id: 1,
		name: "Ethan",
	}

	fmt.Printf("User name: %s\n",user.Name())
}