package main

import "fmt"

func main() {
	fmt.Println(numSub("0110111"))
}

func numSub(s string) int {
	res := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			res = (res + (count+1)*count/2) % 1000000007
			count = 0
		}
	}
	if count > 0 {
		res = (res + (count+1)*count/2) % 1000000007
	}
	return res
}
