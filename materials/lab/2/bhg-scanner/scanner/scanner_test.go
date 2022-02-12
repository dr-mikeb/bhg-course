package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

	totalPorts:= 1024
    open, close := PortScanner(totalPorts) // Currently function returns only number of open ports
	got := open + close
	//gotopen, got close := PortScanner()
    want := 2 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }

}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

	totalPorts := 1024
    gotopen,_ := PortScanner( totalPorts) // Currently function returns only number of open ports
	//open, close repalce got above
	//got := open+close
    want := 1024 // default value; consider what would happen if you parameterize the portscanner ports to scan

    if gotopen != want {
        t.Errorf("got %d, wanted %d", gotopen, want)
    }
}


