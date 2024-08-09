package cards_test

import (
	"horrorsaur/tripletriad/internal/cards"
	"testing"

	f "github.com/brianvoe/gofakeit/v7"
)

func TestCreateCardsWithRanks(t *testing.T) {
	var deck = make([]cards.Card, 10)
	for i := 0; i < 10; i++ {
		deck = append(deck, cards.New(f.Name(), [4]int{4, 5, 7, 9}, cards.Poison, false))
	}
}
