package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"code.justin.tv/spencer/talks/twirptalk"
)

func init() {
	log.SetFlags(0)
}

type streamLogger struct{}

func (l *streamLogger) StartStream(_ context.Context, s *twirptalk.Stream) (*twirptalk.StreamSession, error) {
	id := strconv.Itoa(rand.Intn(1e6))
	log.Printf("starting stream %s with formats %v: id %s", s.ChannelId, s.Formats, id)
	return &twirptalk.StreamSession{Id: id, Stream: s}, nil
}

func (l *streamLogger) EndStream(_ context.Context, s *twirptalk.StreamSession) (*twirptalk.EndStreamResponse, error) {
	log.Printf("ending stream session %s", s.Id)
	return &twirptalk.EndStreamResponse{}, nil
}

func main() {
	s := twirptalk.NewStreamTrackerServer(&streamLogger{}, nil)
	log.Println("launching on port 12345")
	if err := http.ListenAndServe(":12345", s); err != nil {
		log.Fatalf("unable to run: %s", err)
	}
}
