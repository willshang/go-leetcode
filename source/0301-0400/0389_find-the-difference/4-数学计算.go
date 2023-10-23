package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	ch := byte(0)
	for _, value := range t {
		ch += byte(value)
	}
	for _, value := range s {
		ch -= byte(value)
	}
	return ch
}
