package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "5F3Z-2e-9-w"
	K := 4
	fmt.Println(licenseKeyFormatting(S, K))

	S = "2-5g-3-J"
	K = 2
	fmt.Println(licenseKeyFormatting(S, K))
}

func licenseKeyFormatting(S string, K int) string {
	arr := strings.Join(strings.Split(strings.ToUpper(S), "-"), "")
	count := len(arr) / K
	first := len(arr) % K
	if first > 0 {
		count++
	}
	str := arr[:first]
	if first != 0 {
		count = count - 1
	}
	for i := 0; i < count; i++ {
		str = str + "-" + arr[first+i*K:first+(i+1)*K]
	}
	return strings.Trim(str, "-")
}
