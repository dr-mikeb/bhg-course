package scanner

import (
	"testing"
)

func TestFullScan(t *testing.T){

    got := PortScanner() // Currently function returns only number of open ports
    want := 1024 // default value; consider what would happen if you parameterize the portscanner

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}


