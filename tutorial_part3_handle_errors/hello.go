package main

import (
	"fmt"
	"log"
	
)

type GreetingPair struct {
	Greeting string
	RandomGreeting string
}


func Hellos (names []string) (map[string]GreetingPair, error) {

	messages := make(map[string]GreetingPair)

	for thing, name := range names {
		fmt.Println(thing)
		message,err := Hello(name)
		randomMessage, randomErr := RandomHello(name)

		if err != nil {
			return nil, err
		}

		if randomErr != nil {
			return nil, randomErr
		}

		messages[name] = GreetingPair{
			Greeting: message,
			RandomGreeting: randomMessage,
		}


	}
	return messages,nil
}

func main () {
	log.SetPrefix(("greetings: "))
	log.SetFlags(0)

	message,err := Hello("Lol")
	random_message, random_message_err := RandomHello("Wda")

	if err != nil {
		log.Fatal(err)
	}

	if random_message_err != nil {
		log.Fatal(err)
	}

	fmt.Println(random_message)
	fmt.Println(message)

	mapOfGreetingPairs, mapErr := Hellos(
		[]string{
			"Jason",
			"Aaron",
			"Jack",
		},
	)

	if mapErr != nil {
		log.Fatal(mapErr)
	}

	for name, pair := range mapOfGreetingPairs {
		fmt.Printf("User %s", name)
		fmt.Printf("Random Greeting %s", pair.RandomGreeting)
		fmt.Printf("Greeting %s", pair.Greeting)

	}

}