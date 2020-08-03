package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(countDigitOne(13))
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	str := strconv.Itoa(n)
	return dfs(str)
}

func dfs(str string) int {
	if str == "" {
		return 0
	}
	first := int(str[0] - '0')
	if len(str) == 1 && first == 0 {
		return 0
	}
	if len(str) == 1 && first >= 1 {
		return 1
	}
	count := 0
	if first > 1 {
		count = int(math.Pow(float64(10), float64(len(str)-1)))
	} else if first == 1 {
		count, _ = strconv.Atoi(str[1:])
		count = count + 1
	}
	other := first * (len(str) - 1) * int(math.Pow(float64(10), float64(len(str)-2)))
	numLeft := dfs(str[1:])
	return count + numLeft + other
}
