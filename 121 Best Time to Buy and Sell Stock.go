import "math"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	minPrice := math.MaxInt32
	profit := 0
	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		}
		if val-minPrice > profit {
			profit = val - minPrice
		}
	}

	return profit
}
