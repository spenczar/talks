package main

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gogo/protobuf/proto"

	"goji.io/pat"

	"code.justin.tv/spencer/talks/twirptalk"
	goji "goji.io"
)

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
	app := &streamLogger{}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Put("/stream/:name"), func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		var msg *twirptalk.Stream
		if err := proto.Unmarshal(bodyBytes, msg); err != nil {
			w.WriteHeader(400)
			return
		}
		resp, err := app.StartStream(context.Background(), msg)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		respBody, err := proto.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(respBody)
	})
}
