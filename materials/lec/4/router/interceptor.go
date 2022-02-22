// interceptor.go modified from Black Hat Go > CH4 > simple_middleware > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-4/simple_middleware/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// > Potential  go build interceptor.go
// ./interceptor

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("start %s\n", time.Now().String())
	l.Inner.ServeHTTP(w, req)
	log.Printf("finish %s\n", time.Now().String())
}

func hello(w http.ResponseWriter, req *http.Request) {
	mystr := fmt.Sprintf("Hello, we got your request to %s\n", req.URL.Path)
	fmt.Fprint(w, mystr)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: interceptor <port>")
	}
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	myport := fmt.Sprintf(":%s", os.Args[1])
	http.ListenAndServe(myport, &l)
}
