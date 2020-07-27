package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("aababaab"))
}

func removeDuplicates(S string) string {
	arr := []byte(S)
	for {
		flag := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] == arr[i+1] {
				if i+2 != len(arr) {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					arr = arr[:i]
				}
				flag = true
				break
			}
		}
		if flag == false {
			break
		}
	}
	return string(arr)
}
