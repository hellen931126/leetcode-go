import "sort"

var ret [][]int

func help(nums []int, idx int) {
	if idx == len(nums)-1 {
		tmp1 := make([]int, len(nums))
		copy(tmp1, nums)
		ret = append(ret, tmp1)
		return
	}
	sort.Ints(nums[idx:])
	for i := idx; i < len(nums); i++ {
		if i != idx && (nums[i-1] == nums[i]) {
			continue
		}
		tmp := nums[i]
		nums[i] = nums[idx]
		nums[idx] = tmp
		tmp1 := make([]int, len(nums))
		copy(tmp1, nums)
		help(tmp1, idx+1)
		tmp = nums[i]
		nums[i] = nums[idx]
		nums[idx] = tmp
	}
}
func permuteUnique(nums []int) [][]int {
	ret = [][]int{}
	help(nums, 0)
	return ret
}