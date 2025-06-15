## Go Load Test

This repository is for self-learning purposes.  
The goal is to deepen my understanding of Go by implementing a simple load testing tool from scratch.

### Purpose

- Practice Go programming through hands-on development
- Learn about concurrency, HTTP requests, and performance testing

### links
---
locust: https://github.com/locustio/locust  
boomer: https://github.com/myzhan/boomer

### memo
---
### build worker
```
$ go build -o boomer ./boomer/boomer.go
```
### run locust master, worker and local httpServer
```
$ locust --master -f dummy.py
$ go run server/main.go
$ ./boomer/boomer
```

### UI
```
http://localhost:8089/
```
