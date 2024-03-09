package week_0
/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	maxProfit := 0
	lowestPrice := prices[0]
	for _, price := range prices {
		if price < lowestPrice {
			lowestPrice = price
		} else {
			if currentProfit := price - lowestPrice; currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	return maxProfit
}

// @lc code=end
