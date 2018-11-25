//CountWordsAndImages return words and images 
package main

import (
	"fmt"
	http "net/http"
	"bufio"
	"strings"
	"os"
	"log"
	"golang.org/x/net/html"
)

// CountWordsAndImages returns number of words and images in each URL
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return //bare return
		// or
		// return 0, 0, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML :%s", err)
		return // bare
	}
	words, images = countWordsAndImages(doc)
	return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: PROG URL")
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}