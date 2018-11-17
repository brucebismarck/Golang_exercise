// ex3.10 inserts commas into integer strings given as command-line arguments,
// without using recursion.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	postiveIntStart := 0
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		postiveIntStart = 1
	}
	postiveIntEnd := strings.Index(s, ".")
	if postiveIntEnd == -1 {
		postiveIntEnd = len(s)
	}
	newString := s[postiveIntStart:postiveIntEnd]

	pre := len(newString) % 3
	// Write the first group of up to 3 digits.
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(newString[:pre])

	// Deal with the rest.
	for i := pre; i < len(newString); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(newString[i : i+3])
	}
	buf.WriteString(s[postiveIntEnd:])
	return buf.String()
}
