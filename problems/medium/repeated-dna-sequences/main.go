package main

import "fmt"

func main() {
	r := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")

	fmt.Println(r)
}

/*
hash table
	key - hash(sort string) | []values


Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
*/

func findRepeatedDnaSequences(s string) []string {
	store := make(map[string]int)

	for i := 0; i < len(s)-9; i++ {
		currentSeq := s[i : i+10]

		store[currentSeq]++
	}

	var out []string

	for k, v := range store {
		if v > 1 {
			out = append(out, k)
		}
	}

	return out
}
