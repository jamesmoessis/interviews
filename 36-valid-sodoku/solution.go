package main

import "fmt"

func valid(a []byte) bool {
	m := make(map[byte]bool)
	for _, item := range a {
		if _, ok := m[item]; ok && item != '.' {
			return false
		}
		m[item] = true
	}
	return true
}

func solveSquare(square [][]byte) bool {
	a := flatten(square)
	return valid(a)
}

func flatten(square [][]byte) []byte {
	n := len(square)
	a := make([]byte, 0, n*n)
	for _, row := range square {
		for _, item := range row {
			a = append(a, item)
		}
	}
	return a
}

func isValidSudoku(board [][]byte) bool {
	// boxes
	n := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ii, jj := i*3, j*3
			square := board[ii : ii+3]
			square = append(square[:0:0], square...)
			for k := range square {
				square[k] = square[k][jj : jj+3]
			}
			if !solveSquare(square) {
				return false
			}
		}
	}

	// rows
	for _, row := range board {
		if !valid(row) {
			return false
		}
	}

	n = 9
	// columns
	for j := 0; j < 9; j++ {
		column := make([]byte, n)
		for i := 0; i < 9; i++ {
			column[i] = board[i][j]
		}
		if !valid(column) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'1', '2', '3', '.', '.', '.', '4', '5', '6'},
		{'.', '.', '.', '.', '.', '.', '7', '8', '9'},
		{'.', '.', '.', '.', '.', '.', '1', '2', '3'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'1', '2', '3', '.', '.', '.', '4', '5', '6'},
		{'.', '.', '.', '.', '.', '.', '7', '8', '9'},
		{'.', '.', '.', '.', '.', '.', '1', '5', '3'},
	}))
}
