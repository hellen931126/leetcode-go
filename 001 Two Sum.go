func twoSum(nums []int, target int) []int {
	for preKey, preValue := range nums {
		for nextKey := preKey + 1; nextKey < len(nums); nextKey++ {
			if (preValue + nums[nextKey]) == target {
				return []int(preKey, nextKey)
			}
		}
	}

	return []int{}
}
