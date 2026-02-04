package tools

func BinarySearch(list []int, needle int, low, hight int) int {
	for hight >= low {
		middle := (low + hight) / 2

		if list[middle] == needle {
			return middle
		}

		if needle > list[middle] {
			low = middle + 1
		} else {
			hight = middle - 1
		}
	}

	return -1
}

