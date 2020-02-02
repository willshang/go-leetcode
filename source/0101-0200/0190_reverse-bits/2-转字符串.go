package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	str := strconv.FormatUint(uint64(num), 2)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + str[i:i+1]
	}
	if len(rev) < 32 {
		rev = rev + strings.Repeat("0", 32-len(rev))
	}
	n, _ := strconv.ParseUint(rev, 2, 64)
	return uint32(n)
}
