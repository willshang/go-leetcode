package main

import "strings"

func main() {

}

func printVertically(s string) []string {
	arr := strings.Split(s, " ")
	n := len(arr)
	maxLength := 0
	for i := 0; i < n; i++ {
		if len(arr[i]) > maxLength {
			maxLength = len(arr[i])
		}
	}
	res := make([]string, 0)
	for i := 0; i < maxLength; i++ {
		temp := make([]byte, 0)
		for j := 0; j < len(arr); j++ {
			if len(res) < len(arr[j]) {
				temp = append(temp, arr[j][i])
			} else {
				temp = append(temp, ' ')
			}
		}
		res = append(res, strings.TrimRight(string(temp), " "))
	}
	return res
}
