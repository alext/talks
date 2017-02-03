package main

import (
	"bufio"
	"io"
	"os"
	"unicode/utf8"
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

func rot13Byte(in byte) byte {
	return byte(rot13(rune(in)))
}

type rot13Reader struct {
	r *bufio.Reader
}

func NewRot13Reader(r io.Reader) io.Reader {
	return &rot13Reader{
		r: bufio.NewReader(r),
	}
}

// Although this operated on bytes, it's safe to use for unicode streams
// because in UTF-8 all multi-byte characters have the high bit set for all
// bytes. The rot13 function will therefore not change them.
// See https://en.wikipedia.org/wiki/UTF-8#Description
func (r *rot13Reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	for i := 0; i < n; i++ {
		buf[i] = rot13Byte(buf[i])
	}
	return n, err
}

// This implements a reader that operates at a rune level for comparison.
func (reader *rot13Reader) UnicodeRead(buf []byte) (int, error) {
	i := 0
	for {
		r, n, err := reader.r.ReadRune()
		if err != nil {
			return i, err
		}
		if i+n > len(buf) {
			_ = reader.r.UnreadRune()
			if i == 0 {
				// Nothing added to buffer, and the next rune won't fit.
				return 0, io.ErrShortBuffer
			}
			break
		}
		n = utf8.EncodeRune(buf[i:], rot13(r))
		i += n
	}
	return i, nil
}

func rotate(in io.Reader, out io.Writer) {
	reader := NewRot13Reader(in)
	io.Copy(out, reader)
}

func main() {
	rotate(os.Stdin, os.Stdout)
}
