package main

import (
	"fmt"
	"go-game/board"
	"strings"
)

const (
	Size = 20
	MaxRound = 25
)

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
	b := board.CreateBoard(Size)
	fmt.Println("welcome the the best board game")
	fmt.Println("this is your board:")
	fmt.Println(b)
	for round:=1; !b.CheckWin() && round<MaxRound; round++ {  //game loop
		pick := getNumber(board.Keys)
		fmt.Println(fmt.Sprintf("round %v you pick: %v, your new board:", round, pick))
		b.PaintBoard(pick)
		fmt.Println(b)
	}
}
