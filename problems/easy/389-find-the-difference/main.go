package main

import "fmt"

func main() {
	fmt.Println(string(findTheDifference("a", "aa")))
}

func findTheDifference(s string, t string) byte {
	smap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		smap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		smap[t[i]]--
	}

	for i := 0; i < len(t); i++ {
		if count, ok := smap[t[i]]; !ok || (ok && count != 0) {
			return t[i]
		}

	}

	return 0
}
