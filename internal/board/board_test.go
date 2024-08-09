package board

import (
	"testing"
)

func TestBoardGetRow(t *testing.T) {
	board := Board{
		Grid: [3][3]int{
			{1, 1, 1},
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	expected := Row{1, 1, 1}
	if board.GetRow(0) != expected {
		t.Fatalf("expected %x but got %x", board.GetRow(0), expected)
	}
}

func TestBoardGetColumn(t *testing.T) {
	board := Board{
		Grid: [3][3]int{
			{1, 3, 10},
			{1, 8, 8},
			{1, 2, 5},
		},
	}

	expected := [3]int{1, 1, 1}
	if board.GetColumn(0) != expected {
		t.Fatalf("expected %x but got %x", board.GetColumn(0), expected)
	}
}
