package main

import (
	"fmt"
)

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
}

func getHint(secret string, guess string) string {
	length := len(secret)
	right := 0
	wrongLoc := 0
	m := make(map[byte]int)
	n := make(map[byte]int)
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			right++
		} else {
			m[secret[i]]++
			n[guess[i]]++
		}
	}
	for i := range m {
		if m[i] < n[i] {
			wrongLoc = wrongLoc + m[i]
		} else {
			wrongLoc = wrongLoc + n[i]
		}
	}

	return fmt.Sprintf("%dA%dB", right, wrongLoc)
}
