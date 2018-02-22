func pivotIndex(nums []int) int {
	total := sumNums(len(nums), nums)
	for i := 0; i < len(nums); i++ {
		sum := sumNums(i, nums)
		if (sum * 2) == (total - nums[i]) {
			return i
		}
	}
	return -1
}

func sumNums(end int, nums []int) int {
	var sum int
	for i := 0; i < end; i++ {
		sum += nums[i]
	}
	return sum
}