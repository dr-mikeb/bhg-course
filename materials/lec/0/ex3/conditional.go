package main

import "fmt"

func main() {
	// conditional structures
	// - if / else if / else
	test := int(-1)

	if (test == 42) {
		fmt.Println("It's the answer")
	}
	
	
	if test > 0 {
		fmt.Println("greater than zero")
	} else {
		fmt.Println("not greater than zero")	
	}

	if test > 0 || false { // redudant OR --> || false
		fmt.Println("greater than zero")
	} else if test < 0 && true {  // redundant and --> && true
		fmt.Println("less than zero")	
	} else {
		fmt.Println("zero")	
	}

	//switch (as expected)
	schoolString := "uwyo"
	switch schoolString {
		case "uwyo":
			fmt.Println("Go Pokes")
		case "uc":
			fmt.Println("Go Bearcats")
		default:
			fmt.Println("Go Team Go!")
	}

	schoolString = "uc"
	//switch (w/comparisons)
	switch {
	case test == 0:
		fmt.Println("It's zero")
	case test > 0:
		fmt.Println("It's greater than zero")
	case schoolString == "uwyo":
		fmt.Println("Go Pokes")
	default:
		fmt.Println("Less than zero and not the pokes")
	}

	// Repetition

	for i := int(0); i < 10 ; i++{
		fmt.Println(float64(i)/10.0)
	}

	m := make([]string,5)
	m[0] = "Hello"
	m[1] = "my"
	m[2] = "name"
	m[3] = "is"

	for idx, val := range m {
		// modify this statement to remove idx or val; what happens?
		fmt.Printf("The value of m at idx %d is %s\n", idx, val)
	}

	for _, val := range m {	
		fmt.Printf("The value of m is %s\n", val)
	}

	for idx, _ := range m {	
		fmt.Printf("The value of m at idx %d is %s\n", idx, m[idx])
	}

}