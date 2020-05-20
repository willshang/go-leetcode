package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}

func reverseOnlyLetters(S string) string {
	i := 0
	j := len(S) - 1
	arr := []byte(S)
	for i < j {
		if !isLetter(arr[i]) {
			i++
		} else if !isLetter(arr[j]) {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return string(arr)
}

func isLetter(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}
