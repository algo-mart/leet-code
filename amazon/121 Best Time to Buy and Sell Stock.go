package amazon

import "math"

func MaxProfit(prices []int) int {
	max := float64(0)
	i := 0
	for b := 1; i < len(prices) && b < len(prices); b++ {
		if prices[b] > prices[i] {
			if prices[b] < prices[b-1] {
				continue
			}
			max = math.Max(max, float64(prices[b]-prices[i]))
		} else {
			i = b
		}
	}
	return int(max)
}
