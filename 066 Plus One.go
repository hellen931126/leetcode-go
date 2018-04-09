func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] >= 10 {
			digits[i] -= 10
			continue
		}
		return digits
	}
	digits = append([]int{1}, digits...)
	return digits
}