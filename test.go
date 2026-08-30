func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)

	for row := 0; row < numRows; row++ {
		currentRow := make([]int, row+1)

		for col := 0; col <= row; col++ {
			if col == 0 || col == row {
				currentRow[col] = 1
			} else {
				currentRow[col] = result[row-1][col-1] +
					result[row-1][col]
			}
		}

		result = append(result, currentRow)
	}

	return result
}