// bhg-dial.go orginally found in Black Hat Go > CH2 > dial > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/dial/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: go run bhg-dial.go

package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}