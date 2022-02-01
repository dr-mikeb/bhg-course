// bhg-tcp-scanner-wg-too-fast.go modified from Black Hat Go > CH2 > tcp-scanner-wg-too-fast > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-wg-too-fast/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// >Potential  go build -o mywgscan bhg-tcp-scanner-wg-too-fast.go
// > time ./mywgscan > out.csv 

package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("port, status\n")

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()  //decrement wg when completed
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%d, closed/filtered\n", j)
				return
			}
			conn.Close()
			fmt.Printf("%d, open\n", j)
		}(i)
	}
	wg.Wait()  // don't allow the program to terminate until wait group == 0
}