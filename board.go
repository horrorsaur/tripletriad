package main

import (
	"errors"
	"fmt"
)

// Default TT (Triple Triad) Board
//
// https://finalfantasy.fandom.com/wiki/Triple_Triad_(Final_Fantasy_VIII)#Mechanics
const DEFAULT_BOARD_SIZE = 3

type (

	// Triple Triad is played on a three-by-three (3x3) square grid of blank
	// spaces where the cards will be placed.
	Board struct {
		Grid [DEFAULT_BOARD_SIZE][DEFAULT_BOARD_SIZE]int
	}

	Row [DEFAULT_BOARD_SIZE]int
)

func newDefaultBoard() *Board {
	return &Board{Grid: [DEFAULT_BOARD_SIZE][DEFAULT_BOARD_SIZE]int{}}
}

func (b *Board) row(offset int) (Row, error) {
	if offset > DEFAULT_BOARD_SIZE || offset < 0 {
		return Row{0, 0, 0}, errors.New("offset is out of range")
	}
	return b.Grid[offset], nil
}

func (b *Board) update(row, column, value int) {
	b.Grid[row][column] = value
}

// visually print the game board
func (b *Board) print() {
	for i := 0; i < DEFAULT_BOARD_SIZE; i++ {
		for j := 0; j < DEFAULT_BOARD_SIZE; j++ {
			fmt.Print(b.Grid[i][j], " ")
		}
		fmt.Print("\n")
	}
}
