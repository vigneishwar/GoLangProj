package main

import "fmt"

func main() {
	// var card string = "ace of spade"

	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "five of diamonds"
}
