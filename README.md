# Web server implemented in Golang
### To run the program, these are sample commands:
- go run simple-webserver.go -port=80 -listenip=127.0.0.1 -indexfile=./index.htm
- go run simple-webserver.go -port=80 -listenip=127.0.0.1 -indexfile=./index.htm -debug
- go run simple-webserver.go -port=8000 -listenip=127.0.0.1 -indexfile=./index.htm -debug

- Low number ports will require running the program as root.


### Once the web sever is running, use browse the following page:
- http://1ocalhost or http://127.0.0.1 if using port 80
  - or
- http://localhost:8000 or http://127.0.0.1:8000 if using port 8000

