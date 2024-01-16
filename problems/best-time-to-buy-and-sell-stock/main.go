package main

import "fmt"

func main() {
	r := maxProfit([]int{7, 1, 5, 3, 6, 4})
	//						1     1        1      1    3
	// найти все

	fmt.Println(r)
}

func maxProfit(prices []int) int {
	maxProfit, minPrice := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
