// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	fmt.Printf("***********Host Search***********")
	
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h,_ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}


	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}

	
	hostSearchFacets, err := s.HostFacets()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("***********Host Facets***********\n")
	semiformat := fmt.Sprintf("%q\n", *hostSearchFacets)     // Turn the slice into a string that looks like ["one" "two" "three"]
    tokens := strings.Split(semiformat, " ")    // Split this string by spaces
    fmt.Printf(strings.Join(tokens, ", "))
	
	hostSearchFilters, err := s.HostFilters()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("***********Host Filters***********\n")
	semiformat = fmt.Sprintf("%q\n", *hostSearchFilters)     // Turn the slice into a string that looks like ["one" "two" "three"]
    tokens = strings.Split(semiformat, " ")    // Split this string by spaces
    fmt.Printf(strings.Join(tokens, ", ")) //seperate each element with a comma
}