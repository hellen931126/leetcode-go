func maxProfit(k int, prices []int) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}

	if k >= len(prices) {
		return easyMode(prices)
	}

	cur := make([]int, k+1)
	res := make([]int, k+1)

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			cur[j] = max(res[j-1]+max(diff, 0), cur[j]+diff)
			res[j] = max(cur[j], res[j])
		}
	}
	return res[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func easyMode(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}