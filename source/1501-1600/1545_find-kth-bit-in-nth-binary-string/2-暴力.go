package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthBit(4, 11))
}

func findKthBit(n int, k int) byte {
	arr := generate(n)
	return arr[k-1]
}

func generate(n int) []byte {
	arr := make([][]byte, n)
	arr[0] = []byte{'0'}
	for i := 1; i < n; i++ {
		arr[i] = append(arr[i], arr[i-1]...)
		arr[i] = append(arr[i], '1')
		arr[i] = append(arr[i], reverse(invert(arr[i-1]))...)
	}
	return arr[n-1]
}

func reverse(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func invert(arr []byte) []byte {
	for i := 0; i < len(arr); i++ {
		arr[i] = '1' - arr[i] + '0'
	}
	return arr
}
