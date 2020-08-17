package main

import "fmt"

func main() {
	str := "A man, a plan, a canal: Panama"
	reverseString([]byte(str))
	fmt.Println(str)
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
