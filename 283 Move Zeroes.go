func moveZeroes(nums []int) {
	t, s := 0, 0
	for _, i := range nums {
		if i == 0 {
			t += 1
		} else {
			nums[s] = i
			s += 1
		}
	}
	for t > 0 {
		nums[len(nums)-t] = 0
		t -= 1
	}
}