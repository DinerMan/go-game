package main

import (
	"fmt"
	"game/board"
	"strings"
)
const size = 20
const MaxRound = 25

func validInput(pick string, options []string) bool {
	for _, val := range options {
		if pick == val { return true }
	}
	return false
}

func getNumber (options []string) string {
	var pick string
	for !validInput(pick, options){
		fmt.Println("pick a key from the options: ", options)
		fmt.Scanln(&pick); pick = strings.ToUpper(pick)
	}
	return pick
}

func main () {
	// start game
	b := board.CreateBoard(size)
	fmt.Println("welcome the the best board game")
	fmt.Println("this is your board:")
	b.PrintBoard()
	for round:=1; !b.CheckWin() && round<MaxRound; round++ {  //game loop
		pick := getNumber(board.Keys)
		fmt.Println(fmt.Sprintf("round %v you pick: %v, your new board:", round, pick))
		b.PaintBoard(pick)
		b.PrintBoard()
	}
}
