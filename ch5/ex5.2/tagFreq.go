//count the frequency of each tag in a html on stdin
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func tagFreq(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0) // create a dictionary
	z := html.NewTokenizer(os.Stdin)
	var err error
	for {
		type_ := z.Next()
		if type_ == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++ //If key exist then add one. Because the dict in go start with nil/0, no need to have else.
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}

func main() {
	freq, err := tagFreq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}
