package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// START OMIT
var input = strings.NewReader("The Go Programming Language")

func main() {
	reader := bufio.NewReader(input)
	for {
		word, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading STDIN:", err)
		}
		fmt.Printf("Read string '%s'\n", word)
	}
}

// END OMIT
