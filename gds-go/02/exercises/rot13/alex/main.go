package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rot13(in rune) rune {
	if (in >= 'a' && in <= 'm') || (in >= 'A' && in <= 'M') {
		return in + 13
	} else if (in >= 'n' && in <= 'z') || (in >= 'N' && in <= 'Z') {
		return in - 13
	} else {
		return in
	}
}

func rotate(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading STDIN:", err)
		}

		fmt.Fprint(out, string(rot13(r)))
	}
}

func main() {
	rotate(os.Stdin, os.Stdout)
}
