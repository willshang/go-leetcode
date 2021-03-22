package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minCharacters("aba", "caa"))
	fmt.Println(minCharacters("dabadd", "cda"))
}

// leetcode1737_满足三条件之一需改变的最少字符数
func minCharacters(a string, b string) int {
	arrA, arrB := [26]int{}, [26]int{}
	for i := 0; i < len(a); i++ {
		arrA[a[i]-'a']++
	}
	for i := 0; i < len(b); i++ {
		arrB[b[i]-'a']++
	}
	res := math.MaxInt32
	lengthA, lengthB := len(a), len(b)
	sumA, sumB := 0, 0
	// a-y的情况
	for i := 0; i < 25; i++ {
		sumA = sumA + arrA[i]
		sumB = sumB + arrB[i]
		res = min(res, lengthA-sumA+sumB)               // 条件1：a<b
		res = min(res, sumA+lengthB-sumB)               // 条件2: b<a
		res = min(res, lengthA-arrA[i]+lengthB-arrB[i]) // 条件3：全都相同
	}
	res = min(res, lengthA-arrA[25]+lengthB-arrB[25]) // z全都相同
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
