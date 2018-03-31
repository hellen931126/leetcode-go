func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length-1] {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}