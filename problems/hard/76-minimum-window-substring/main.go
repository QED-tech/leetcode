package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))

}

/*
 ABC -> func contains(substring) -> true/false

       r
  l
 ADOBECODEBANC

---- windowStore ----
 A - 0
 D - 1
 O - 1
 B - 1
 E - 1
 C - 1
--- END windowStore ----

---- tStore ----
 tStore for  ABC
 A - 1
 B - 1
 C - 1
---- END tStore ----

1. Match ? windowStore[l]-- ; l++ : default

----- CODE BLOCK ----
 left = 0
 for r := 1; r < len(s); r++ {
 	for contains(l..r) && l < r {
		l++
		min = min(min, r - l + 1)
	}
 }
----- END CODE BLOCK ----

*/

func minWindow(s string, t string) string {
	tStore := make(map[string]int)
	windowStore := make(map[string]int)

	for _, r := range t {
		tStore[string(r)]++
	}

	var (
		l       int
		minimum = len(s) + 1
		ans     string
	)

	for r := 0; r < len(s); r++ {
		windowStore[string(s[r])]++

		for match(windowStore, tStore) {
			windowStore[string(s[l])]--

			if r-l+1 < minimum {
				minimum = r - l + 1
				ans = s[l : r+1]
			}

			l++
		}

	}

	return ans
}

func match(window map[string]int, tStore map[string]int) bool {
	for el, count := range tStore {
		wCount := window[el]
		if !(wCount >= count) {
			return false
		}
	}

	return true
}
