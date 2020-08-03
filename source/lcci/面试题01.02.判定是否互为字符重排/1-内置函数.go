package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}

func CheckPermutation(s1 string, s2 string) bool {
	arr1 := strings.Split(s1, "")
	arr2 := strings.Split(s2, "")
	sort.Strings(arr1)
	sort.Strings(arr2)
	return strings.Join(arr1, "") == strings.Join(arr2, "")
	// return reflect.DeepEqual(arr1, arr2)
}
