package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	str = strings.Replace(str, "6", "9", 1)
	res, _ := strconv.Atoi(string(str))
	return res
}
