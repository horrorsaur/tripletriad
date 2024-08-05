package board

import (
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	board := newDefaultBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.update(i, j, 1)
		}
	}

	board.print()

	row, err := board.row(0)
	if err != nil {
		t.Fatal(err)
	}

	if Sum(row) != 3 {
		t.Fatalf("expected row sum to be 3 but got %d", Sum(row))
	}
}
