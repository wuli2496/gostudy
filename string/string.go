package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	left := n % 3
	for i := 0; i < n; i++ {
		if i >= left && 0 == (i - left) % 3 {
			buf.WriteByte(',')
		}

		buf.WriteByte(s[i])
	}

	return buf.String()

}
func main() {
	s := "abcde"

	fmt.Println(comma(s))
}