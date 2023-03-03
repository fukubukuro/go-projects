package main

import (
	"fmt"
	. "tic-tac-go/models"
	. "tic-tac-go/services"
)

func main() {
	board := Board{}
	printBoard(&board)
	piece := X
	for true {
		fmt.Println("it's", piece.String()+"'s", "turn")
		outcome := processTurn(&board, piece)
		if outcome != "" {
			fmt.Println(outcome)
			break
		}
		printBoard(&board)
		piece = switchTurns(piece)
	}
}

func processTurn(board *Board, piece Piece) string {
	place, err := ReadPlace()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if 0 == place {
		printHelp()
		return ""
	}
	err = board.Place(piece, place)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return board.Outcome()
}

func printBoard(board *Board) {
	printRow(board[0])
	printSeparator()
	printRow(board[1])
	printSeparator()
	printRow(board[2])
	fmt.Println("type a whole number between 1 and 9 or h for help")
}

func printRow(row Row) {
	fmt.Println("", row[0], "|", row[1], "|", row[2])
}

func printSeparator() {
	fmt.Println("---|---|---")
}

func printHelp() {
	fmt.Println("", 1, "|", 2, "|", 3)
	printSeparator()
	fmt.Println("", 4, "|", 5, "|", 6)
	printSeparator()
	fmt.Println("", 7, "|", 8, "|", 9)
	fmt.Println("place your piece according to this layout")
	fmt.Println("===========")
}

func switchTurns(piece Piece) Piece {
	if X == piece {
		return O
	} else {
		return X
	}
}
