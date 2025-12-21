package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(1))
}

	func countAndSay(n int) string {
		if n == 0 {
			return ""
		}

		var (
			counter int
		)

		var fn func(c int, compact string) string
		fn = func(c int, compact string) string {
			if c+1 == n {
				return compact
			}

			return fn(c+1, compactString(compact))
		}

		return fn(counter, "1")
	}

	func compactString(orig string) string {
		if orig == "" {
			return ""
		}

		var (
			b       strings.Builder
			counter = 1
		)

		for i := 0; i < len(orig); i++ {
			curr := orig[i]

			if i+1 >= len(orig) {
				b.WriteString(fmt.Sprintf("%d%s", counter, string(curr)))
				break
			}

			if !(curr == orig[i+1]) {
				b.WriteString(fmt.Sprintf("%d%s", counter, string(curr)))

				counter = 1
			} else {
				counter++
			}
		}

		return b.String()
	}
