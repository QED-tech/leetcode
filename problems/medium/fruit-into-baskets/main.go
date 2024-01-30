package main

import "fmt"

/*

Example 1:

Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.
Example 2:

Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].
Example 3:

Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].

*/

func main() {
	// fmt.Println(totalFruit([]int{1, 2, 1}))
	// fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

/*
    |l|                   |r|
1,   1,  1,    3,    2,    2


store = {
	3:1
	2:1
}

if len(store) > 2 {
	for len(store) > 2 {
		del()
		left++
	}
}

*/

func totalFruit(fruits []int) int {
	store := make(map[int]int)
	var left, ans int

	for right := 0; right < len(fruits); right++ {
		store[fruits[right]]++
		if len(store) > 2 {
			for len(store) > 2 {
				store[fruits[left]]--
				if store[fruits[left]] == 0 {
					delete(store, fruits[left])
				}

				left++
			}
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
