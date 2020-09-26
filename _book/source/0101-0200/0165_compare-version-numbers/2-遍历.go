package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.1", "1.10"))
	fmt.Println(compareVersion("1.03", "1.02001"))
}

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
		a := strings.TrimLeft(arr1[i], "0")
		b := strings.TrimLeft(arr2[i], "0")
		// 1 < 10 => 01 < 10
		for len(a) < len(b) {
			a = "0" + a
		}
		for len(b) < len(a) {
			b = "0" + b
		}
		for j := 0; j < len(a); j++ {
			if a[j] < b[j] {
				return -1
			} else if a[j] > b[j] {
				return 1
			}
		}
	}
	return 0
}
