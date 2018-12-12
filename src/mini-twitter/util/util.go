package util

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var mux sync.Mutex

func RaftAPICall(ctx context.Context, method string, url string, payload *strings.Reader, responseChannel chan *http.Response) {
	mux.Lock()

	log.Println("url =", url)
	log.Println("method =", method)
	log.Println("payload =", payload)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println("Error received while creating  =", err)
		mux.Unlock()
		return
	}

	select {
	case <-time.After(100 * time.Nanosecond):
		log.Println("Time exceeded so calling Raft server", url)
	case <-ctx.Done():
		log.Println(ctx.Err())
		mux.Unlock()
		return

	}
	log.Println("Process called", url)

	var res *http.Response
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error received from Raft =", err)
		mux.Unlock()
		return
	}

	log.Println("res =", res)
	mux.Unlock()

	responseChannel <- res

	log.Println("after channel is finished =", url)
}

func InteractWithRaftStorage(method string, key string, value interface{}) (string, error) {

	log.Println("Interacted with Raft, method called =", method)
	var payloadValue string
	var res *http.Response
	if method != "GET" {
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(value); err != nil {
			log.Println("Error occured while encoding ", key, " data =", err)
			return "", err
		}
		payloadValue = buf.String()
	}

	responseChannel := make(chan *http.Response)

	ports := [3]string{"12380", "22380", "32380"}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	for _, port := range ports {
		var payload *strings.Reader
		payload = nil
		if value != nil {
			payload = strings.NewReader(payloadValue)
		}
		url := "http://127.0.0.1:" + port + "/" + key

		go RaftAPICall(ctx, method, url, payload, responseChannel)

	}

	select {
	case r := <-responseChannel:
		log.Println("r =", r)
		res = r
		cancel()

	case <-time.After(5 * time.Second):
		log.Println("Raft Server timedout")
		return "", errors.New("Raft server timedout!!")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error occured while decoding response from Raft =", err)
		return "", err
	}

	log.Println("data received from Raft after calling ", method, " method =", string(data))

	return string(data), nil
}
