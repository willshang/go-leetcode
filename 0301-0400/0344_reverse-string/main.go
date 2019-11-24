package main

import "fmt"

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(reverseString(str))
}
func reverseString(s string) string {
	bytes := []byte(s)

	i, j := 0, len(s)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
