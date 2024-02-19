package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "G()(al)"
	fmt.Println(interpret(str))
	fmt.Println(interpret("G()()()()(al)"))
	fmt.Println(interpret("(al)G(al)()()G"))
	fmt.Println(interpret(""))
}

/*
Example 1:

Input: command = "G()(al)"
Output: "Goal"
Explanation: The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".
Example 2:

Input: command = "G()()()()(al)"
Output: "Gooooal"
Example 3:

Input: command = "(al)G(al)()()G"
Output: "alGalooG"
*/

func interpret(command string) string {
	var sb strings.Builder

	for i := 0; i < len(command); i++ {
		if string(command[i]) == "G" {
			sb.WriteString("G")
		} else if string(command[i]) == "(" && string(command[i+1]) == ")" {
			sb.WriteString("o")
		} else if string(command[i]) == "(" && string(command[i+1]) == "a" {
			sb.WriteString("al")
		}
	}

	return sb.String()
}
