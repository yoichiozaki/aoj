package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	board := make([][]string, 4)
	for i := range board {
		scanner.Scan()
		board[i] = strings.Split(scanner.Text(), " ")
	}
	board = rotate90(board)
	board = rotate90(board)
	for i := range board {
		fmt.Println(strings.Join(board[i], " "))
	}
}

func rotate90(board [][]string) [][]string {
	rotated := make([][]string, 4)
	for i := range rotated {
		rotated[i] = make([]string, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rotated[j][3-i] = board[i][j]
		}
	}
	return rotated
}
