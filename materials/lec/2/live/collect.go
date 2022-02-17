package main

import (
	"fmt"
)

func main(){
	arr := []rune{"a", "b", "c", "d", "e"}

	// Use a for loop with manual index itterator
	for i := 0; i < len(arr); i++ { 

	}

	// Use a for loop with index/val itterator
	for index, val := range arr  { 

	}
	
	// Use a for loop with index/_ itterator
	for index, _ := range arr  { 

	}

	// Use a for loop with index/_ itterator
	for _, val := range arr  { 

	}

	

}