package main

import "fmt"

// setZeroes sets an entire row and column to zero when a cell is zero.
// Time: O(rows*cols), Extra space: O(1)
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	rows, cols := len(matrix), len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// Use the first row and first column as markers.
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

	if firstColZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)

	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}