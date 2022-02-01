// bhg-copy-example.go modified from Black Hat Go > CH2 > copy-example > main.go and io-example > main.go

// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/copy-example/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// >Potential  go run bhg-copy-example.go

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)
// GopherReader defines an io.Reader to read from stdin.
type HumanReader struct{}
// Read reads data from stdin.
func (humanReader *HumanReader) Read(b []byte) (int, error) {
	fmt.Print("human input > ")
	return os.Stdin.Read(b)
}
// SpecialWriter defines an io.Writer to write to Stdout.
type GopherWriter struct{}
// Write writes data to Stdout.
func (gopherWriter *GopherWriter) Write(b []byte) (int, error) {
	fmt.Print("excited gopher says > ")
	res1 := bytes.ReplaceAll(b, []byte(" "), []byte("!"))
	return os.Stdout.Write(res1)
}
func main() {
	// Instantiate reader and writer
	var (
		reader HumanReader
		writer GopherWriter
	)

	input := make([]byte, 4096)
	humansays, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from the human on stdin\n",humansays)

	gophersays, err := writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("As gopher I wrote %d bytes to stdout\n",gophersays)


	// Instead of doing this manually 
	// We could just use the io.Copy that takes a reader and writer; 
	// abstracted out: reader calls the Read function while writer calls the Write function
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}