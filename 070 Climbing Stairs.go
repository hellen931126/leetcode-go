func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	ways := make([]int, n+1)
	ways[0], ways[1] = 1, 1

	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}