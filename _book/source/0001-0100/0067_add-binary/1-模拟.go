package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	length := len(a)

	A := transToInt(a, length)
	B := transToInt(b, length)

	return makeString(add(A, B))
}

func transToInt(s string, length int) []int {
	result := make([]int, length)
	ls := len(s)
	for i, b := range s {
		result[length-ls+i] = int(b - '0')
	}
	return result
}

func add(a, b []int) []int {
	length := len(a) + 1
	result := make([]int, length)
	for i := length - 1; i >= 1; i-- {
		temp := result[i] + a[i-1] + b[i-1]
		result[i] = temp % 2
		result[i-1] = temp / 2
	}
	i := 0
	for i < length-1 && result[i] == 0 {
		i++
	}
	return result[i:]
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}
