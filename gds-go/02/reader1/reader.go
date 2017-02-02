package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var input = strings.NewReader("Hello")

// START OMIT
func main() {
	reader := bufio.NewReader(input)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading STDIN:", err)
		}
		fmt.Printf("Read a byte: %x, in ascii: %c\n", b, b)
	}
}

// END OMIT
