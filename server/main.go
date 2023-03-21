package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func ping(w http.ResponseWriter, r *http.Request) {
	latency := time.Duration(rand.Intn(1050)) // timeout when latency >= 1000
	time.Sleep(latency * time.Millisecond)
	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", ping)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
