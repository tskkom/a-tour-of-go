package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fs := strings.Fields(s)
	m := make(map[string]int)
	for _, w := range fs {
		v, ok := m[w]
		if ok {
			m[w] = v + 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
