package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}

/*func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.Trim(s," ")," ")
	return len(arr[len(arr)-1])
}*/

func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	result := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result != 0 {
				return result
			}
			continue
		}
		result++
	}
	return result
}
