package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; i >= 0 && j >= 0 && k >= 0; k-- {
		if nums2[j] >= nums1[i] {
			nums1[k] = nums2[j]
			j--
			continue
		}
		nums1[k] = nums1[i]
		i--
	}
	for ; j >= 0; j-- {
		nums1[j] = nums2[j]
	}
	return
}
