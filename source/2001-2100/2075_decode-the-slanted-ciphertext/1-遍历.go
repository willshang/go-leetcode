package main

import "strings"

func main() {

}

// leetcode2075_解码斜向换位密码
func decodeCiphertext(encodedText string, rows int) string {
	a, b := rows, len(encodedText)/rows
	res := make([]byte, 0)
	for i := 0; i < b; i++ {
		x, y := 0, i
		for x < a && y < b {
			res = append(res, encodedText[x*b+y])
			x++
			y++
		}
	}
	return strings.TrimRight(string(res), " ")
}
