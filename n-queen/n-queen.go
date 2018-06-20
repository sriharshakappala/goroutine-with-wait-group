package main

import "fmt"

const n = 8

var board [n][n]int

func printBoard(board [n][n]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d", board[i][j])
		}
		fmt.Println()
	}
}

func isSafe(board [n][n]int, row int, col int) bool {
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; j >= 0 && i < n; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func solveNQueen(board *[n][n]int, col int) bool {
	if col >= 8 {
		printBoard(*board)
		return true
	}
	for i := 0; i < n; i++ {
		if isSafe(*board, i, col) {
			board[i][col] = 1
			if solveNQueen(board, col+1) {
				return true
			}
			board[i][col] = 0
		}
	}
	return false
}

func main() {
	if solveNQueen(&board, 0) == false {
		fmt.Println(false)
	}
	fmt.Println(true)
	printBoard(board)
}
