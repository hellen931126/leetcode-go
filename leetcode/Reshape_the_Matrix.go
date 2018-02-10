func matrixReshape(nums [][]int, r int, c int) [][]int {
	height := len(nums)
	width := len(nums[0])
	if height*width != r*c {
		return nums
	}

	matrix := make([][]int, r)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, c)
	}
	idx := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			matrix[idx/c][idx%c] = num[i][j]
			idx++
		}
	}
	return matrix
}