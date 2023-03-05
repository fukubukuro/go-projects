package main

import (
	"fmt"
	. "tic-tac-go/models"
	. "tic-tac-go/services"
)

func main() {
	board := Board{}
	gameOver := false
	piece := X
	printBoard(&board)
	for true {
		printPrompt()
		fmt.Println("it's", piece.String()+"'s", "turn:")
		gameOver, piece = processTurn(&board, piece)
		printBoard(&board)
		if gameOver {
			break
		}
	}
}

func processTurn(board *Board, piece Piece) (bool, Piece) {
	place, err := ReadPlace()
	fmt.Println("")
	if err != nil {
		fmt.Println(err)
		return false, piece
	}
	if 0 == place {
		printHelp()
		return false, piece
	}
	err = board.Place(piece, place)
	if err != nil {
		fmt.Println(err)
		return false, piece
	}
	outcome := board.Outcome()
	if outcome != "" {
		fmt.Println(outcome)
		return true, piece
	} else {
		return false, switchTurns(piece)
	}
}

func printBoard(board *Board) {
	printRow(board[0])
	printSeparator()
	printRow(board[1])
	printSeparator()
	printRow(board[2])
}

func printPrompt() {
	fmt.Println("")
	fmt.Println("type a whole number between 1 and 9 or h for help")
	fmt.Println("")
}

func printRow(row Row) {
	fmt.Println("", row[0], "|", row[1], "|", row[2])
}

func printSeparator() {
	fmt.Println("---|---|---")
}

func printHelp() {
	fmt.Println("")
	fmt.Println("place your piece according to this layout:")
	fmt.Println("")
	fmt.Println("", 1, "|", 2, "|", 3)
	printSeparator()
	fmt.Println("", 4, "|", 5, "|", 6)
	printSeparator()
	fmt.Println("", 7, "|", 8, "|", 9)
	fmt.Println("")
	fmt.Println("==========================================")
	fmt.Println("")
}

func switchTurns(piece Piece) Piece {
	if X == piece {
		return O
	} else {
		return X
	}
}
