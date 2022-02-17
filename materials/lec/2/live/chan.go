package main

import (
	"fmt"
)

func SendDataToChannel(ch chan int, value int) {
    ch <- value
}
 
func main() {
 
    var reciever int
    intchannel := make(chan int)     

    go SendDataToChannel(intchannel, 2022)   // send data w/ go routine


    reciever = <-intchannel                    // receive data 
 
    fmt.Println(reciever)  
}