package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var input = strings.NewReader("Héllo ☃")

// START OMIT
func main() {
	reader := bufio.NewReader(input)
	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading STDIN:", err)
		}
		fmt.Printf("Read %d bytes rune: %U, character: %c\n", n, r, r)
	}
}

// END OMIT
