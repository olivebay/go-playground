/*

You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an error.

*/

package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second


// Racer compares the response times of a and b, returning the fastest one, timing out on 10s
func Racer(a, b string) (wunner string, error error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func measureResponseTime(url string, timeout time.Duration) (duration time.Duration, error error) {
	start := time.Now()
	client := http.Client{Timeout: timeout}
	_, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}

// ConfigurableRacer compares the response times of a and b, returning the fastest one
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-getURL(a):
		return a, nil
	case <-getURL(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out while waiting for %q and %q", a, b)
	}
}

func getURL(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
