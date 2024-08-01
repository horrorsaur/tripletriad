package main

// Pink cards belong to the opponent and blue cards belong to the player

type (
	Card struct {
		Name  string
		Ranks [4]Rank
	}

	// Each card has four numbers (known as ranks).
	//
	// Each number corresponds to one of the four sides of the card. The ranks range from one to ten.
	Rank int // 10 == A

	// The top right corner of the card sometimes has an elemental symbol representing the card's element:
	// Earth, Fire, Water, Poison, Holy, Lightning, Wind, or Ice
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

var elementNames = map[Element]string{
	Fire: "Fire",
}
