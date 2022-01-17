package main

import (
	"fmt"
)

func strlen(s string, c chan int){
	fmt.Printf("Working on String %s\n",s)
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Go!", c)
	go strlen("Pokes", c)
	go strlen("!Go", c)
	go strlen("Cowboys", c)

	x, y, z, w := <-c, <-c, <-c, <-c 

	fmt.Printf("Total length is %d\n", x+w+y+z)
}

