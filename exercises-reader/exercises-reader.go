package main

import "golang.org/x/tour/reader"

type MyReader struct {}

// a Reader type that emits an infinite stream of the ASCII character 'A'.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
