package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"code.justin.tv/spencer/talks/twirptalk"
)

func init() {
	log.SetFlags(0)
}

func main() {
	c := twirptalk.NewStreamTrackerProtobufClient("http://127.0.0.1:12345", http.DefaultClient)

	everySecond(func(idx int) {
		resp, err := c.StartStream(context.Background(), &twirptalk.Stream{
			ChannelId: fmt.Sprintf("%d", idx),
			Formats:   []string{"high", "medium", "low", "mobile", "source"},
		})
		if err != nil {
			log.Fatalf("err: %s", err)
		}
		log.Printf("stream %d started, session id=%s", idx, resp.Id)
	})
}

func everySecond(f func(i int)) {
	var i = 0
	for range time.NewTicker(time.Second).C {
		f(i)
		i++
	}
}
