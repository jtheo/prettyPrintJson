package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// dont do this, see above edit
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return b, err
	}
	return out.Bytes(), nil
}

func main() {
	var err error
	var fname string
	body := make([]byte, 2048)
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	input := os.Stdin
	if len(os.Args) == 2 {
		fname = os.Args[1]
		input, err = os.Open(fname)
		if err != nil {
			log.Fatalf("opening the file %s, got an error: %v\n", fname, err)
		}

		defer input.Close()
	}

	lenFile, err := input.Read(body)
	if verbose {
		fmt.Printf("I've read %d bytes from %s\n", lenFile, fname)
	}

	if err != nil {
		log.Fatalf("reading the file %s of len %d, got an error: %v\n", fname, lenFile, err)
	}

	b, err := prettyprint(body[:lenFile])
	if err != nil {
		log.Printf("prettyprint returned an error: %v\n\n", err)
		if verbose {
			log.Printf("%x", b)
		}
	}

	fmt.Printf("%s\n", b)
}
