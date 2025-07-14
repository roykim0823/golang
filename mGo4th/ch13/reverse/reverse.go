package main

// This version has bugs

import (
	"fmt"
)

func R1(s string) []byte {
	sAr := []byte(s)
	rev := make([]byte, len(sAr))

	l := len(sAr)
	for i := 0; i < l; i++ {
		rev[i] = sAr[l-1-i]
	}

	return rev
}

func R2(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func main() {
	str := "1234567890"
	fmt.Println(string(R1(str)))
	fmt.Println(R2(str))
}
