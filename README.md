for study.  
write load test scenarios in Go.

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

UI
```
http://localhost:8089/
```
