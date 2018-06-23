package main

import "fmt"

const n = 8

var board [n][n]int
var row = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
var column = [8]int{1, 2, 2, 1, -1, -2, -2, -1}

func printBoard(board [n][n]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func isValid(x int, y int) bool {
	if x < 0 || y < 0 || x >= n || y >= n {
		return false
	}
	return true
}

func knightTour(board [n][n]int, x int, y int, position int) {
	board[x][y] = position
	if position >= n*n {
		fmt.Println("--------")
		printBoard(board)
	}
	for i := 0; i < n; i++ {
		newX := x + row[i]
		newY := y + column[i]
		if (isValid(newX, newY)) && (board[newX][newY] == 0) {
			knightTour(board, newX, newY, position+1)
		}
	}
	board[x][y] = 0
}

func main() {
	position := 1
	knightTour(board, 0, 0, position)
	fmt.Println("Over")
}
