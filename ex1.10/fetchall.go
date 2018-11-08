//Fetchall fetches URLs in parallel and reports their times and sizes
package main

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)//channel
	for _, url := range os.Args[1:]{
		go fetch(url, ch) // start a go routine
	}
	for range os.Args[1:]{
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string){
	start := time.Now()
	resp, err:= http.Get(url)
	if err!= nil {
		ch <- fmt.Sprint(err) //send to channel ch
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // copy from back to front. Dont need to read into memory
	resp.Body.Close() //dont leak resources

	if err!= nil{
		ch <- fmt.Sprintf("while reading %s :%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}