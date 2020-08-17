package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.00000000000000000000000000000000000000000000000000000000101", "1.00100"))
}

// leetcode165_比较版本号
func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	for len(arr1) < len(arr2) {
		arr1 = append(arr1, "0")
	}
	for len(arr2) < len(arr1) {
		arr2 = append(arr2, "0")
	}
	for i := 0; i < len(arr1); i++ {
		a, _ := strconv.Atoi(arr1[i])
		b, _ := strconv.Atoi(arr2[i])
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}
