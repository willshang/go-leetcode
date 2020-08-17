package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

// leetcode1108_IP地址无效化
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
