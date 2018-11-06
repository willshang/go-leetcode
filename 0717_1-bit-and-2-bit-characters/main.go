package main

import "fmt"

func main() {
	arr := []int{1,0,0}
	fmt.Println(isOneBitCharacter(arr))
}
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1{
		if bits[i] == 1{
			i = i + 2
		}else {
			i++
		}
	}
	return i == n-1
}