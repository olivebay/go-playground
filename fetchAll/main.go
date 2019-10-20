// Fetch URLs specified as arguments concurrently and print the elapsed time and size of each one
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Exiting.. enter URLs to fetch")
		os.Exit(1)
	}

	start := time.Now()

	ch := make(chan string)

	fmt.Println("Fetching..")

	// for each URL create a go routine and fetch each url
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	// Print elapsed time
	fmt.Printf("%.2fs total elapsed\n", time.Since(start).Seconds())
}

// fetch each url and return the size and elapsed time of each fetch
// Output: 0.55s     217 https://httpbin.org/get
func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get("http://" + url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %s", err) // send to channel ch
		return
	}
	defer resp.Body.Close() //dont leak resources

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// Return the size and elapsed time of each fetch
	ch <- fmt.Sprintf("%.2fs %7d bytes %s", secs, nbytes, url)
}
