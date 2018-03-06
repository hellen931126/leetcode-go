import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) <= 2 {
		return ret
	}
	sort.Ints(nums)
	if nums[0]+nums[1] > 0 {
		return ret
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}
