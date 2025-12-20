package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

/*
[-4 -1 1 2]

|i|     |l|             |r|
-4      -1      1       2
sum = -3, diff = 4

|i|             |l|     |r|
-4      -1      1       2
sum = -1, diff = 2

        |i|     |l|     |r|
-4      -1      1       2
sum = 2, diff = 1


[1 2 3 4 5 10 13 20 22 30 33]
|i|     |l|                                                                     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 36, diff = 4

|i|             |l|                                                             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 37, diff = 3

|i|                     |l|                                                     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 38, diff = 2

|i|                             |l|                                             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 39, diff = 1

|i|                                     |l|                                     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 44, diff = 4

|i|                                     |l|                             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 41, diff = 1

|i|                                     |l|                     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 33, diff = 7

|i|                                             |l|             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 36, diff = 4

|i|                                                     |l|     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 43, diff = 3

        |i|     |l|                                                             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 38, diff = 2

        |i|             |l|                                                     |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 39, diff = 1

        |i|                     |l|                                             |r|
1       2       3       4       5       10      13      20      22      30      33
sum = 40, diff = 0
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var (
		closest = nums[0] + nums[1] + nums[2]
	)

	for i := 0; i < len(nums)-2; i++ {

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}

			if diffrence(sum, target) < diffrence(closest, target) {
				closest = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}

func diffrence(sum, target int) int {
	return int(math.Abs(float64(target - sum)))
}
