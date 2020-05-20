package main

import "fmt"

func main() {
	arr := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(arr))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 1, 0}))
}

func isOneBitCharacter(bits []int) bool {
	return helper(bits, 0)
}

func helper(bits []int, left int) bool {
	if left == len(bits)-1 {
		return bits[left] == 0
	}
	if left < len(bits)-1 {
		if bits[left] == 0 {
			return helper(bits, left+1)
		}
		if bits[left] == 1 {
			return helper(bits, left+2)
		}
	}
	return false
}
