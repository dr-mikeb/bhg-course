// bhg-slow-scanner.go modified from Black Hat Go > CH2 > tcp-slow-scanner > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-slow/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// >Potential  go build -o mysscan bhg-slow-scanner.go
// > time ./mysscan > out.csv 


package main

// DRMIKE: Adding in time for DialTimeout 
import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Printf("port, status\n")  // adding output - csv format
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second ) // DRMIKE: using DialTimeout over Dial (one extra parameter!); consider < 1 Second
		if err != nil {
			fmt.Printf("%d, closed/filtered\n", i)  // adding output - csv format
			continue
		}
		conn.Close()
		fmt.Printf("%d, open\n", i) // modified output for csv format
	}
}