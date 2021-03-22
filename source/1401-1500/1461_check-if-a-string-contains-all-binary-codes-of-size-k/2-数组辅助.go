package main

import "fmt"

func main() {
	fmt.Println(hasAllCodes("00110", 2))
}

func hasAllCodes(s string, k int) bool {
	length := 1 << k
	arr := make([]bool, length)
	cur := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		cur = cur<<1 + num
		if i >= k-1 {
			cur = cur & (length - 1)
			arr[cur] = true
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == false {
			return false
		}
	}
	return true
}
