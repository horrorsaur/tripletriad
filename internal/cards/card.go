package cards

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type (
	GameDeck struct {
		Cards []Card
	}

	Card struct {
		Name string `json:"name"`

		// Each card has four numbers (known as ranks).
		//
		// Each number corresponds to one of the four sides of the card. The ranks range from one to ten.
		Ranks [4]int `json:"ranks"` // 10 == A

		// The top right corner of the card sometimes has an elemental symbol representing the card's Element:
		// Earth, Fire, Water, Poison, Holy, Lightning, Wind, or Ice
		Element element `json:"element"`
	}

	element int

	// describes the souped json to encode
	classicCardGroup struct {
		Names []string `json:"names"`
		Type  string   `json:"type"`
		Level string   `json:"level"`
	}
)

const (
	None element = iota
	Fire
	Earth
	Water
	Poison
	Holy
	Lightning
	Wind
	Ice
)

var elements = map[element]string{
	Fire:      "fire",
	Earth:     "earth",
	Water:     "water",
	Poison:    "poison",
	Holy:      "holy",
	Lightning: "lightning",
	Wind:      "wind",
	Ice:       "ice",
}

var DefaultDeck GameDeck

// Generates the original FF8 Triple Triad cards and stores the output in the DefaultDeck
func Initialize() bool {
	cards := generateDefaultTTCards()
	DefaultDeck.Cards = cards
	return true
}

func New(name string, ranks [4]int, element element, shuffle bool) Card {
	return Card{
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

func setRankByLevel(_ string) [4]int {
	var ranks [4]int
	for i := 0; i < 4; i++ {
		n := rand.Intn(9)
		if n > 0 {
			ranks[i] = n
		}
		ranks[i] = n + 1 // temp
	}
	return ranks
}

func generateDefaultTTCards() []Card {
	cardByts, _ := os.ReadFile("../../data/original_card_names.json")

	var data []classicCardGroup
	if err := json.Unmarshal(cardByts, &data); err != nil {
		log.Fatal(err)
	}

	var generatedCards []Card // len == 88
	for _, card := range data {
		for _, name := range card.Names {
			card := New(name, setRankByLevel(card.Level), None, false)
			generatedCards = append(generatedCards, card)
		}
	}

	return generatedCards
}
