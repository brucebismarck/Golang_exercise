//Dup2 prints the count and tesxt of lines that appear more than
// once in the standard input. It reads from stdin or from a lst of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil { //zero value for pointers, interfaces,
				//maps, slices, channels and function types, representing an uninitialized value.
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, foundIn[line], line)
			//map[replace2:[replace.txt] replace3:[replace.txt] test:[test.txt]
			// test1:[test.txt] test2:[test.txt] test3:[test.txt] replace:[replace.txt] replace1:[replace.txt]]
			// this is foundIn

		}
	}
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), foundIn[line]) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
		//if counts[line] > 1 {
		//	fmt.Printf("%v\n", f.Name())
		//}
	}

}
