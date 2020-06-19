package main

import (
	"fmt"
	"net/http"
	"time"
)

func pageSize(url string) (size int64, err error) {
	// pageSize returns the size in bytes of the HTTP response for url
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	return resp.ContentLength, nil
}

func main() {
	start := time.Now()
	dmtns := make(chan string)

	for id := 1; id < 150; id++ {
		url := fmt.Sprintf("https://dmtn-%03d.lsst.io", id)
		go func() {
			size, err := pageSize(url)
			if err != nil {
				dmtns <- fmt.Sprintf("err retrieving %s: %v", url, err)
			} else {
				dmtns <- fmt.Sprintf("%s: %d", url, size)
			}
		}()
	}

	for msg := range dmtns {
		fmt.Printf("%s: %s\n", time.Since(start), msg)
	}
}
