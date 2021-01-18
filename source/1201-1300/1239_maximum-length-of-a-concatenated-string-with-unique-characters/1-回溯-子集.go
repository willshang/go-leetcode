package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}

var res []string

func maxLength(arr []string) int {
	res = make([]string, 0)
	dfs(arr, "", 0)
	maxValue := 0
	for i := 0; i < len(res); i++ {
		if check(res[i]) == true && len(res[i]) > maxValue {
			maxValue = len(res[i])
		}
	}
	return maxValue
}

func dfs(arr []string, path string, index int) {
	res = append(res, path)
	for i := index; i < len(arr); i++ {
		dfs(arr, path+arr[i], i+1)
	}
}

func check(s string) bool {
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		value := int(s[i] - 'a')
		arr[value]++
		if arr[value] > 1 {
			return false
		}
	}
	return true
}
