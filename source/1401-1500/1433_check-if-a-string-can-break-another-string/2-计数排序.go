package main

import (
	"fmt"
)

func main() {
	//fmt.Println(checkIfCanBreak("abc", "xya"))
	fmt.Println(checkIfCanBreak("leetcodee", "interview"))
	//fmt.Println(checkIfCanBreak("abe", "acd"))
}

func checkIfCanBreak(s1 string, s2 string) bool {
	arr1 := [26]int{}
	arr2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		arr1[int(s1[i]-'a')]++
		arr2[int(s2[i]-'a')]++
	}
	a, b := 0, 0
	totalA, totalB := 0, 0
	for i := 0; i < 26; i++ {
		totalA = totalA + arr1[i]
		totalB = totalB + arr2[i]
		if totalA >= totalB {
			a++
		}
		if totalB >= totalA {
			b++
		}
	}
	return a == 26 || b == 26
}
