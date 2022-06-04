package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func discountPrices(sentence string, discount int) string {
	arr := strings.Fields(sentence)
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == '$' && len(arr[i]) > 1 &&
			strings.Count(arr[i], "$") == 1 && check(arr[i][1:]) == true {
			v, _ := strconv.Atoi(arr[i][1:])
			vf := float64(v) * float64(100-discount) / 100
			arr[i] = fmt.Sprintf("$%0.2f", vf)
		}
	}
	return strings.Join(arr, " ")
}

func check(s string) bool {
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			continue
		}
		return false
	}
	return true
}
