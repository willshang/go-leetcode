package main

import "strconv"

func main() {

}

// leetcode2243_计算字符串的数字和
func digitSum(s string, k int) string {
	arr := []byte(s)
	for len(arr) > k {
		temp := make([]byte, 0)
		for i := 0; i < len(arr); i = i + k {
			sum := 0
			for j := i; j < i+k && j < len(arr); j++ {
				sum = sum + int(arr[j]-'0')
			}
			temp = append(temp, []byte(strconv.Itoa(sum))...)
		}
		arr = temp
	}
	return string(arr)
}
