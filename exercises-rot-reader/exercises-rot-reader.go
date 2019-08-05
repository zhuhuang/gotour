package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	bytes := make([]byte, len(b))
	n, err := r13.r.Read(bytes)

	for i := range bytes {
		b[i] = rot13(bytes[i])
	}

	return n, err
}

func rot13(b byte) byte {
	if unicode.IsUpper(rune(b)) {
		return 'A' + (b - 'A' + 13) % 26
	} else if unicode.IsLower(rune(b)) {
		return 'a' + (b - 'a' + 13) % 26
	} else {
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
