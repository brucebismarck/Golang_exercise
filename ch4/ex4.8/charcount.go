// Charcount computes counts of Unicode characters.
package main

import (
	"bufio" //Package bufio implements buffered I/O
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // count of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	letterct := 0
	digitct := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //returns rune, nbytes, error
		if err == io.EOF {         // End of file, the only error will be EOF
			break // Because this is the only error, we can stop the loop
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 { //rune == ReplacementChar
			invalid++
			continue
		}
		if unicode.IsLetter(r) == true {
			letterct++
			continue
		}
		if unicode.IsDigit(r) == true {
			digitct++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if letterct > 0 {
		fmt.Printf("\n%d letters\n", letterct)
	}
	if digitct > 0 {
		fmt.Printf("\n%d digits\n", digitct)
	}
}
