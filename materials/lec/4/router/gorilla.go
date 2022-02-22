// gorilla.go modified from Black Hat Go > CH4 > Text Page 82
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// > Potential  go build gorilla.go
// ./gorilla

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type logger struct {
	Inner http.Handler
}

func main() {
	//Modified based on example on pp. 82
	r := mux.NewRouter()
	r.HandleFunc("/class/{studentname}", func(w http.ResponseWriter, req *http.Request){
		name := mux.Vars(req)["studentname"]
		fmt.Fprintf(w ,"I see you %s\n",name)
	}).Methods("GET")
	//Modified based on example on pp. 83
	r.HandleFunc("/strictclass/{studentname:[a-z]+[0-9]*}", func(w http.ResponseWriter, req *http.Request){
		name := mux.Vars(req)["studentname"]
		fmt.Fprintf(w ,"I see you %s\n",name)
	})
	// more at https://github.com/gorilla/mux > readme 
	r.HandleFunc("/varclass/{class:[a-Z]+}/{studentname:[a-z]+[0-9]*}", func(w http.ResponseWriter, req *http.Request){
		classtitle := mux.Vars(req)["class"]
		name := mux.Vars(req)["studentname"]
		fmt.Fprintf(w ,"%s I see you in %s\n",name,classtitle)
	})
	log.Fatal(http.ListenAndServe(":4010", r))
}
