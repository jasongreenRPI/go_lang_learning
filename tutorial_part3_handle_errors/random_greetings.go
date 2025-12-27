package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func RandomHello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message,nil


}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Welcome %v!, good to see you",
		"Hail %v!",
	}

	return formats[rand.Intn(len(formats))]
}