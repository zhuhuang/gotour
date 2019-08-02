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
	for {
		_, err := r13.r.Read(bytes)
		if err == io.EOF {
			break
		}
	}

	for i := range bytes {
		if unicode.IsUpper(rune(bytes[i])) {
			b[i] = 'A' + (bytes[i] - 'A' + 13) % 26
		} else if unicode.IsLower(rune(bytes[i])) {
			b[i] = 'a' + (bytes[i] - 'a' + 13) % 26
		} else {
			b[i] =  bytes[i]
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
