// simple-router.go modified from Black Hat Go > CH4 > simple-router > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-4/simple-scanner/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// > Potential  go build simple-router.go
// ./simple-router

package main

import (
	"fmt"
	"net/http"
	"math/rand" // Dr.Mike Adding for more fun examples
	"time" // Dr.Mike Adding for more fun examples
)


type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	
	rand.Seed(time.Now().UnixNano())
	const maxint = int(^uint(0)>>1) // https://stackoverflow.com/a/6878625

	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a")
	case "/b":
		fmt.Fprint(w, "Executing /b")
	case "/c":
		fmt.Fprint(w, "Executing /c")

	case "/index.html":
		fmt.Fprint(w, "<html><head><title>Real Website I Promise</title></head><body><H1>Welcome</H1>Nothing to see here</body></html>")
	case "/r":
		fmt.Fprint(w, rand.Intn(maxint))
	case "/essay":
		bytes := make([]byte, 1000)
		for i := 0; i < 1000; i++ {
			bytes[i] = byte(rand.Intn(122-97))+97
		}
		fmt.Fprint(w, string(bytes))

	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}