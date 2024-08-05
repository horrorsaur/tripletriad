package cards

import (
	"fmt"
	"math/rand"
)

type (
	Card struct {
		Name string `json:"name"`

		// Each card has four numbers (known as ranks).
		//
		// Each number corresponds to one of the four sides of the card. The ranks range from one to ten.
		Ranks [4]int `json:"ranks"` // 10 == A

		// The top right corner of the card sometimes has an elemental symbol representing the card's Element:
		// Earth, Fire, Water, Poison, Holy, Lightning, Wind, or Ice
		Element Element `json:"element"`
	}

	Element int
)

const (
	Fire Element = iota
	Earth
	Water
	Poison
	Holy
	Lightning
	Wind
	Ice
)

var elements = map[Element]string{
	Fire:      "fire",
	Earth:     "earth",
	Water:     "water",
	Poison:    "poison",
	Holy:      "holy",
	Lightning: "lightning",
	Wind:      "wind",
	Ice:       "ice",
}

func New(name string, ranks [4]int, element Element, shuffle bool) *Card {
	return &Card{
		Name:    name,
		Ranks:   handleRanks(ranks, shuffle),
		Element: element,
	}
}

func handleRanks(ranks [4]int, shuffle bool) [4]int {
	for i := 0; i < len(ranks); i++ {
		if ranks[i] > 10 || ranks[i] <= 0 {
			panic(fmt.Sprintf("rank at pos %d exceeds max rule. got %d", i, ranks[i]))
		}
	}

	if shuffle {
		rand.Shuffle(len(ranks), func(i, j int) {
			ranks[i], ranks[j] = ranks[j], ranks[i]
		})
	}
	return ranks
}
