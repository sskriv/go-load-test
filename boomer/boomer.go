package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/myzhan/boomer"
)

const (
	URL_PING = "http://localhost:8080/ping"
)

func worker() {
	request, err := http.NewRequest(http.MethodPost, URL_PING, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := &http.Client{
		Timeout: time.Second,
	}

	startTime := time.Now()
	response, err := client.Do(request)
	elapsed := time.Since(startTime)

	if err != nil {
		boomer.RecordFailure("http", "error", 0.0, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("read resp error: %v\n", err)
	} else {
		log.Printf("Status Code: %d\n", response.StatusCode)
		log.Println(string(body))
	}

	boomer.RecordSuccess(
		"http",
		"worker",
		elapsed.Nanoseconds()/int64(time.Millisecond),
		response.ContentLength,
	)
}

func main() {
	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     worker,
	}

	boomer.Run(task)
}
