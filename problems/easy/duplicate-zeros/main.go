package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(arr)

	fmt.Println(arr)
}

/*
	1, 0, 2, 3, 0, 4, 5, 0
    1  0  0  2  3  0   4  5


out=   1,0,0,2,3,0,0,4

*/

func duplicateZeros(arr []int) {
	for left := 0; left < len(arr)-1; left++ {
		if arr[left] == 0 {
			elem := arr[left+1]
			for right := left + 2; right < len(arr); right++ {
				tmp := arr[right]
				arr[right] = elem
				elem = tmp
			}
			arr[left+1] = 0

			left++
		}
	}
}
