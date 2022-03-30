package scanner

import (
	"testing"
	// "fmt"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){
	
    got, _ := PortScanner(20,81) // Currently function returns only number of open ports
    
	want := []int{22, 80} // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan		  
	for i, val:= range got{
		if val!=want[i]{
			t.Errorf("got %d, wanted %d", got[i], want[i])
		}
	}
	
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()
    got1, got2 := PortScanner(20,81) // Currently function returns only number of open ports
	//got := open + closed
    want1 := []int{22,80} // default value; consider what would happen if you parameterize the portscanner ports to scan
	want2 := []int{20,21,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,
		49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,81}
	
	for i:= 0; i<len(got1); i++{
		if got1[i]!=want1[i]{
			t.Errorf("Open Port: got %d, wanted %d at position %d", got1[i], want1[i], i)
		}
	}
	for i:= 0; i<len(got2); i++{
		if got2[i]!=want2[i]{
			t.Errorf("Closed Port: got %d, wanted %d at position %d", got2[i], want2[i], i)
		}
	}
}


