package board

import (
	"fmt"
	"math/rand"
	"reflect"
)


const color = "\033["
const colorReset = "0m"
var keyMap = map[string]int{
	"G": 32,
	"R": 31,
	"Y": 33,
	"B": 34,
}

var Keys = func () []string {
	keys := reflect.ValueOf(keyMap).MapKeys()
	var formatKeys []string
	for _, key := range keys {
		formatKeys = append(formatKeys, key.String())
	}
	return formatKeys
}()


type GameBoard struct {
	board [][]string
}

func CreateBoard(size int) GameBoard {
	//declare new board
	newBoard := make([][]string, size, size)
	ln := len(Keys)  //make sure that you only call len once
	// initial values
	for i := 0; i < size; i++ {
		newRow := make([]string, size, size)
		for y := 0; y < size; y++ {
			newRow[y] = Keys[rand.Intn(ln)]
		}
		newBoard[i] = newRow
	}
	return GameBoard{board: newBoard}
}

func (g GameBoard) PrintBoard() {
	for _, row := range g.board {
		for _, value := range row {
			fmt.Print(color ,keyMap[value], "m", value, " ")
		}
		fmt.Println(color, colorReset)
	}
}

func (g GameBoard) CheckWin() bool {
	needed_color := g.board[0][0]
	for _, row := range g.board {
		for _, col := range row {
			if col != needed_color {
				return false
			}
		}
	}
	return true
}

func paintKey(cache [][]bool, board [][]string, oldKey string, newKey string, ln int, x int, y int) {
	if x<0 || x>=ln || y<0 || y>=ln { return }
	if cache[x][y] { return }
	// never run on thought's coordinate
	if board[x][y] == oldKey {
		board[x][y] = newKey
		paintKey(cache, board, oldKey, newKey, ln, x, y-1)
		paintKey(cache, board, oldKey, newKey, ln, x, y+1)
		paintKey(cache, board, oldKey, newKey, ln, x-1, y)
		paintKey(cache, board, oldKey, newKey, ln, x+1, y)
	}
}

func (g GameBoard) PaintBoard(newKey string) {
	oldKey := g.board[0][0]
	cache := make([][]bool, len(g.board))
	for index, _ := range cache { cache[index] = make([]bool, len(g.board[0])) }
	paintKey(cache, g.board, oldKey, newKey, len(cache), 0, 0)
}