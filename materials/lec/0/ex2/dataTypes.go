package main

import "fmt"

func main() {
	// Go Syntax: Primitative Basics
	z := int(42)
	var y = int(42)
	var x = "Hello Uwyo!"
	w := "I'm a string"

	fmt.Printf("z = %d, y = %d, x = %s, w=%s\n", z,y,x,w)

	// Go Sytanx: Complex Built-in Slices/Maps
	var s = make([]string, 0)
	var m = make(map[string]int)

	s = append(s, "Go")
	s = append(s, "Pokes")
	m["answer"] = 42
	m["a"] = 0

	fmt.Println("s is ", s)
	fmt.Println("m is ",m)

	// Data Types: Complex - Pointers

	var count = int(42)
	ptr := &count
	fmt.Printf("The memory address of count (&count) = %d\n", &count)
	fmt.Printf("Variable ptr (ptr) contents are a memory address = %d\n", ptr)
	fmt.Println("Variable ptr points (*ptr) to a memory address that holds value:", *ptr)

	*ptr = 100
	fmt.Printf("Modified the variable at memory location (%d) which was count, now = %d \n", ptr, count)
	
}