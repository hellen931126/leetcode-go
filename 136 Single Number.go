func singleNumber(nums []int) int {
	re := 0
	for _, n := range nums {
		re ^= n
	}
	return re
}