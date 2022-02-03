// bhg-slow-scanner.go modified from Black Hat Go > CH2 > tcp-slow-scanner > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-slow/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// >Potential  go build -o myfscan bhg-tcp-scanner-too-fast.go
// > time ./myfscan > out.csv 

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("port, status\n")  // adding output - csv format

	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%d, closed/filtered\n", i)  // adding output - csv format
				return
			}
			conn.Close()
			fmt.Printf("%d, open\n", j)
		}(i)
	}
}