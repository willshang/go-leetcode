package main

import (
	"fmt"
)

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
}

// leetcode299_猜数字游戏
func getHint(secret string, guess string) string {
	length := len(secret)
	right := 0
	wrongNum := 0
	m := make(map[int]int)
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			right++
		}
		m[int(secret[i]-'0')]++
		m[int(guess[i]-'0')]--
	}
	for i := range m {
		if m[i] > 0 {
			wrongNum = wrongNum + m[i]
		}
	}
	// wrongLoc = 总数 - 猜对的数 - 猜错的数
	wrongLoc := length - right - wrongNum
	return fmt.Sprintf("%dA%dB", right, wrongLoc)
}
