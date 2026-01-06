package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}

/*
i - 4
love - 3
leetcode - 2
coding - 2
foo - 1

1. create map
2. create arr [2,1,2,3,4]
3. sort arr [4,3,2,2,1]
4. take k elems from arr[:k]
5.
*/

type state struct {
	word  string
	count int
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	var sl []state
	for w, c := range m {
		sl = append(sl, state{
			word:  w,
			count: c,
		})
	}

	sort.Slice(sl, func(i, j int) bool {
		if sl[i].count != sl[j].count {
			return sl[i].count > sl[j].count
		}

		return sl[i].word < sl[j].word
	})

	var out []string

	for i := range k {
		out = append(out, sl[i].word)
	}

	return out
}
