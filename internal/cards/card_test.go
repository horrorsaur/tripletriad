package cards_test

import (
	"horrorsaur/tripletriad/internal/cards"
	"testing"
)

func TestSanityCheckDefaultDeck(t *testing.T) {
	ok := cards.Initialize()
	if len(cards.DefaultDeck.Cards) < 88 || !ok {
		t.Fatal("unexpected default cards len ", len(cards.DefaultDeck.Cards))
	}
}
