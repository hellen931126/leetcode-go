func removeElement(nums []int, val int) int {
	non_val_size := 0
	for _, i := range nums {
		if i != val {
			nums[non_val_size] = i
			non_val_size += 1
		}
	}
	return non_val_size
}