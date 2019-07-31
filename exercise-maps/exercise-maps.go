package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	chars := strings.Fields(s)
	for _, char := range chars {
		element, present := m[char]
		if present {
			m[char] = element + 1
		} else {
			m[char] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}