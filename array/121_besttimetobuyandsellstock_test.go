package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
[(3) Best Time to Buy and Sell Stock - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
*/
func TestBestTimeToBuyAndSellStock(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	expect := 5
	ans := maxProfit(input)
	assert.Equal(t, expect, ans)
}

func maxProfit(prices []int) int {
	s, f, profit := 0, 0, 0
	for f < len(prices) {
		if prices[f]-prices[s] < 0 {
			s = f
		}
		profit = max(prices[f]-prices[s], profit)
		f++
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
