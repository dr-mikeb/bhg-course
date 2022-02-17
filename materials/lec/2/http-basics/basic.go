package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func newRequest(verb string, targetURL string) {
	fmt.Printf("\n******Attempting a %s request to %s\n",verb,targetURL)
	req, err := http.NewRequest(verb, targetURL, nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Printf("Ending %s request\n",verb)
	} 

func main() {
	primaryURL := "https://www.google.com/robots.txt"

	fmt.Printf("\n******Attempting a %s request to %s\n","GET",primaryURL)

	resp, err := http.Get(primaryURL)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("\n******GET RESPONSE - STATUS %s\n",primaryURL)
	// Print HTTP Status
	fmt.Println(resp.Status)

	fmt.Printf("\n******GET RESPONSE - BODY \n")
	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
	fmt.Printf("Ending %s request\n","GET")	
	
	fmt.Printf("\n******Attempting a %s request to %s\n","GHEAD","primaryURL")
		resp, err = http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Printf("Ending %s request\n","HEAD")

	fmt.Printf("\n******Attempting a %s request to %s\n","POST - FORM",primaryURL)
	form := url.Values{}
	form.Add("foo", "bar")
	resp, err = http.Post(
		"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Printf("Ending %s request\n","POST - FORM")
	
	// Wrapped New Requestn	ewRequest	("DELETE", primaryURL)
	newRequest("DELETE", primaryURL)	
	newRequest("PUT", primaryURL)	
	newRequest("GET", primaryURL)	

}