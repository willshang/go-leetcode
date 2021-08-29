package main

import "fmt"

func main() {
	fmt.Println(minTimeToType("bza"))
}

func minTimeToType(word string) int {
	res := 0
	cur := 0
	for i := 0; i < len(word); i++ {
		target := abs(int(word[i]-'a') - cur)
		res = res + 1 + min(target, 26-target)
		cur = int(word[i] - 'a')
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
