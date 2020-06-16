package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	//fmt.Println(isLongPressedName("alex", "alexxr"))
	//fmt.Println(isLongPressedName("vtkgn", "vttkgnnggg"))
	//fmt.Println(isLongPressedName("vtkgn", "vttkgnngg"))
}

// leetcode925_长按键入
func isLongPressedName(name string, typed string) bool {
	i := 0
	j := 0
	for j < len(typed) {
		if i == len(name) {
			i = len(name) - 1
		}
		if name[i] == typed[j] {
			// 正确的话，保证i == len(name) && j == len(typed)
			i++
			j++
		} else {
			if i == 0 {
				return false
			}
			if name[i-1] != typed[j] {
				return false
			} else {
				j++
			}
		}
	}
	return i == len(name) && j == len(typed)
}
