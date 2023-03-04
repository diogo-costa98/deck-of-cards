package main

import "github.com/diogo-costa98/deck-of-cards/api"

func main() {
	err := api.StartServer()
	if err != nil {
		panic(err)
	}
}
