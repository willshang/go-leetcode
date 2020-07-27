package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	sum := 0
	mul := 1
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		value := int(str[i] - '0')
		sum = sum + value
		mul = mul * value
	}
	return mul - sum
}
