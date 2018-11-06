package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("AA"))
}
func titleToNumber(s string) int {
	result := 0

	for i := 0; i < len(s); i++{
		temp := int(s[i] - 'A' + 1)
		result = result * 26 + temp
	}
	return result
}