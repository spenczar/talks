//+build ignore
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var _ = sync.WaitGroup{}

func size(body io.ReadCloser) int {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return len(bytes)
}

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

	for id := 1; id < 150; id++ {
		url := fmt.Sprintf("https://dmtn-%03d.lsst.io", id)
		size, err := pageSize(url)
		if err != nil {
			fmt.Printf("err retrieving %s: %v\n", url, err)
		}
		fmt.Printf("%s: %s: %d\n", time.Since(start), url, size)
	}
}

// END OMIT
