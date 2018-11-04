// Comments
package main

import (
	"fmt"
	"os" // here after "fmt", we must add ; to split the line.
)

// " " and ' ' are two different thigns

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // no indentation requests
		// := is short variable indentation, so you dont need to declare what is i in the var part.
		// {} can be anywhere. No need to start another line
		s += sep + os.Args[i]
		sep = " " // After first one, add space to seperate two arguments
		fmt.Println(i)
		fmt.Println(os.Args[i])
	}
	//fmt.Println(os.Args[0]) //First argument is code location.
	// /var/folders/g3/x409mz5n43z6g613gky722380000gn/T/go-build011047887/b001/exe/echo1
	fmt.Println(s)
}
