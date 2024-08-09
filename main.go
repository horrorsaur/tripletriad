package main

import (
	"flag"
	"fmt"
	"horrorsaur/tripletriad/internal/cards"
)

var (
	boardSize = flag.Int("size", 3, "sets the triple triad boards size. the default is 3. the maxiumum is 5 (experimental)")
)

func main() {
	flag.Parse()

	if ok := cards.Initialize(); ok {
		fmt.Println("default cards generated")
	}

	fmt.Println(cards.DefaultDeck.Cards[0:8])
}
