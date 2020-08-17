package main

import "fmt"

func main() {
	fmt.Println(oneEditAway("pale", "ple"))
}

func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	i := 0
	for ; i < len(first) && i < len(second); i++ {
		if first[i] != second[i] {
			if len(first) == len(second) {
				if first[i+1:] == second[i+1:] {
					return true
				}
			} else if len(first) < len(second) {
				if first[i:] == second[i+1:] {
					return true
				}
			} else {
				if first[i+1:] == second[i:] {
					return true
				}
			}
			break
		}
	}
	if i == len(first) || i == len(second) {
		return true
	}
	return false
}
