package main

import "fmt"

// NewReverse reverses a slice of ints in place
func reverse(ints *[5]int) *[5]int { // restricted length
	for i := 0; i < len(ints)/2; i++ { // because len(ints) is a integer, i < 2
		end := len(ints) - 1 - i
		ints[i], ints[end] = ints[end], ints[i]
	}
	return ints
}

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reverse(&s))
}
