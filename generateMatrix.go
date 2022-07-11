func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if n == 0 {
		return matrix
	}

	rowStart := 0
	rowEnd := n - 1
	colStart := 0
	colEnd := n - 1
	num := 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			matrix[rowStart][i] = num
			num++
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = num
			num++
		}
		colEnd--
		for i := colEnd; i >= colStart; i-- {
			if rowStart <= rowEnd {
				matrix[rowEnd][i] = num
				num++
			}
		}
		rowEnd--
		for i := rowEnd; i >= rowStart; i-- {
			if colStart <= colEnd {
				matrix[i][colStart] = num
				num++
			}
		}
		colStart++
	}
	return matrix
}
