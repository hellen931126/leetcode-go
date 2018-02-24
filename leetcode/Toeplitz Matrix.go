/*
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
*/

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 1 {
		return true
	}
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

