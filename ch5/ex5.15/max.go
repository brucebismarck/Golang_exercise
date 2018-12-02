package main

import "fmt"

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals {
		if m < v {
			m = v
		}
	}
	return m
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals {
		if m > v {
			m = v
		}
	}
	return m
}

func max2(first int, vals ...int) int {
	m := first
	for _, v := range vals {
		if m < v {
			m = v
		}
	}
	return m
}

func min2(first int, vals ...int) int {
	m := first
	for _, v := range vals {
		if m > v {
			m = v
		}
	}
	return m
}

func main() {
	fmt.Println(min(3, -1, 4))
	fmt.Println(max(3, -1, 4))
	fmt.Println(min2(3, -1, 4))
	fmt.Println(max2(3, -1, 4))
}
