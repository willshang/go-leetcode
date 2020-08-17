package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(hammingWeight(15))
}

func hammingWeight(num uint32) int {
	return strings.Count(strconv.FormatInt(int64(num), 2), "1")
	// return strings.Count(fmt.Sprintf("%b",num),"1")
}
