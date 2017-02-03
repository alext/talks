package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func streamAddition(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	var numbers []int
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading STDIN:", err)
		}

		if b >= '0' && b <= '9' {
			num := int(b - '0')
			numbers = append(numbers, num)
			if len(numbers) == 3 {
				fmt.Fprint(out, (numbers[0]+numbers[1]+numbers[2])%10)
				numbers = numbers[1:] // Remove leading element
			}
		}
	}
}

func main() {
	streamAddition(os.Stdin, os.Stdout)
	fmt.Print("\n")
}
