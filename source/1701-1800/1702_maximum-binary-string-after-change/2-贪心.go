package main

import "strings"

func main() {

}

func maximumBinaryString(binary string) string {
	n := len(binary)
	count := strings.Count(binary, "1")
	if count >= n-1 {
		return binary
	}
	indexZero := strings.IndexByte(binary, '0')
	count = count - indexZero
	return strings.Repeat("1", n-1-count) + "0" + strings.Repeat("1", count)
}
