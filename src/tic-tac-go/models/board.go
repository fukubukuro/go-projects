package models

import "errors"

type Piece int

func (p Piece) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

const (
	EMPTY Piece = iota
	X
	O
)

type Row [3]Piece
type Board [3]Row

type Coords struct {
	X int
	Y int
}

var errInvalidPlace = errors.New("invalid place: provide a whole number between 1 and 9")
var errOccupiedPlace = errors.New("invalid place: already occupied")

func (b *Board) Place(p Piece, place int) error {
	if place < 1 || place > 9 {
		return errInvalidPlace
	}
	coords := intToCoords(place)
	if b[coords.X][coords.Y] != EMPTY {
		return errOccupiedPlace
	}
	b[coords.X][coords.Y] = p
	return nil
}

func intToCoords(i int) Coords {
	i -= 1
	return Coords{i / 3, i % 3}
}

func (board *Board) Outcome() string {
	winningPositions := [8][3]int{}
	winningPositions[0] = [3]int{1, 2, 3}
	winningPositions[1] = [3]int{4, 5, 6}
	winningPositions[2] = [3]int{7, 8, 9}
	winningPositions[3] = [3]int{1, 4, 7}
	winningPositions[4] = [3]int{2, 5, 8}
	winningPositions[5] = [3]int{3, 6, 9}
	winningPositions[6] = [3]int{1, 5, 9}
	winningPositions[7] = [3]int{3, 5, 7}
	for w := 0; w < 8; w++ {
		winner := board.won(winningPositions[w])
		if winner != EMPTY {
			return winner.String() + " wins!"
		}
	}
	filled := true
	for n := 1; n <= 9; n++ {
		place := intToCoords(n)
		if filled {
			filled = board[place.X][place.Y] != EMPTY
		}
	}
	if filled {
		return "cats game."
	}
	return ""
}

func wonRow(row *Row) Piece {
	current := EMPTY
	for n := 0; n < 3; n++ {
		piece := row[n]
		if piece == EMPTY || piece != current {
			return EMPTY
		}
		current = piece
	}
	return current
}

func (board *Board) won(place [3]int) Piece {
	current := EMPTY
	for n := 0; n < 3; n++ {
		coords := intToCoords(place[n])
		piece := board[coords.X][coords.Y]
		if piece == EMPTY {
			return EMPTY
		}
		if current != EMPTY && current != piece {
			return EMPTY
		}
		current = piece
	}
	return current
}
