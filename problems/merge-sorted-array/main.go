package main

import "fmt"

// 1, 2,

func main() {
	a := []int{1, 2, 3, 30, 40, 50, 0, 0, 0, 0}
	b := []int{5, 10, 35, 45}

	//[5 30 30 35 40 45 50]

	merge(a, 6, b, 4)

	fmt.Println(a)
}

func merge(leftList []int, leftItemsNum int, rightList []int, rightItemsNum int) {
	endPointer := leftItemsNum + rightItemsNum - 1

	leftItemsNum--
	rightItemsNum--

	for rightItemsNum >= 0 {
		if leftItemsNum >= 0 && leftList[leftItemsNum] >= rightList[rightItemsNum] {
			leftList[endPointer] = leftList[leftItemsNum]
			leftItemsNum--
		} else {
			leftList[endPointer] = rightList[rightItemsNum]
			rightItemsNum--
		}

		endPointer--
	}
}
