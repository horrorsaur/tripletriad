package board

// Default TT (Triple Triad) Board
//
// https://finalfantasy.fandom.com/wiki/Triple_Triad_(Final_Fantasy_VIII)#Mechanics
const DEFAULT_BOARD_SIZE int = 3

type (

	// Triple Triad is played on a three-by-three (3x3) square grid of blank
	// spaces where the cards will be placed.
	Board struct {
		Grid [DEFAULT_BOARD_SIZE][DEFAULT_BOARD_SIZE]int
	}

	Row [DEFAULT_BOARD_SIZE]int
)

func NewDefaultBoard() *Board {
	return &Board{Grid: [DEFAULT_BOARD_SIZE][DEFAULT_BOARD_SIZE]int{}}
}

func (b Board) GetRow(offset int) Row {
	if offset < 0 || offset >= DEFAULT_BOARD_SIZE {
		return Row{0, 0, 0} // out of bounds, no need for error
	}
	row := b.Grid[offset]
	return Row{row[0], row[1], row[2]}
}

func (b Board) GetColumn(offset int) [3]int {
	if offset < 0 || offset >= DEFAULT_BOARD_SIZE {
		return [3]int{0, 0, 0} // out of bounds, no need for error
	}

	var col [3]int
	for i := 0; i < len(b.Grid); i++ {
		col[i] = b.Grid[i][offset]
	}
	return col
}

func (b *Board) update(row, column, value int) {
	b.Grid[row][column] = value
}

func (r Row) Sum() int {
	var tmp int
	for _, v := range r {
		tmp = tmp + v
	}
	return tmp
}
