//Function to eliminate duplicates in a []string slice
package main

import "fmt"

func dedup(strs []string) []string {
	w := 0 // index of last written string
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}

func main() {
	s := []string{"a", "a", "b", "c", "d"}
	fmt.Println(dedup(s))
}
