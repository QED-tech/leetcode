package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

/*

e a t
t e a
a t e
aet = [eat, tea, ate]

*/

func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)

	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)

		m[string(b)] = append(m[string(b)], s)
	}

	var res [][]string
	for _, el := range m {
		res = append(res, el)
	}

	return res
}
