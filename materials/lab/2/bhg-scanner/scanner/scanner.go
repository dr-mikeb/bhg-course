// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int  // notice the capitalization here. access limited!


func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.Dial("tcp", address) // TODO 2 : REPLACE THIS WITH TIMEOUT (before testing!)
		if err != nil { 
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider making PortScanner take a variable for the ports to scan (int? slice? ); make pass in a target address?
func PortScanner() int {  

	ports := make(chan int, 100)   // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	return len(openports) // TODO 6 : Return total number of ports scanned
}
