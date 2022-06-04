package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

// leetcode2288_价格减免
func discountPrices(sentence string, discount int) string {
	arr := strings.Fields(sentence)
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == '$' {
			v, err := strconv.Atoi(arr[i][1:])
			if err == nil {
				vf := float64(v) * float64(100-discount) / 100
				arr[i] = fmt.Sprintf("$%0.2f", vf)
			}
		}
	}
	return strings.Join(arr, " ")
}
