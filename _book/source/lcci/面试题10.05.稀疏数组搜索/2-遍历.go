package main

import "fmt"

func main() {
	fmt.Println(findString([]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta"))
}

func findString(words []string, s string) int {
	for i := 0; i < len(words); i++ {
		if s == words[i] {
			return i
		}
	}
	return -1
}
