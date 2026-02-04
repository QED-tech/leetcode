package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteString("#")
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	var (
		out []string
		i   = 0
	)

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		j++

		out = append(out, encoded[j:j+length])
		i = j + length
	}

	return out
}

func main() {
	s := Solution{}

	encoded := s.Encode([]string{"hello", "world"})

	fmt.Println(encoded)
	fmt.Println(s.Decode(encoded))
}
