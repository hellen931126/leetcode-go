func missingNumber(nums []int) int {

	miss := 0

	for i := 0; i < len(nums); i++ {

		miss ^= (i + 1) ^ nums[i]

	}

	return miss

}
