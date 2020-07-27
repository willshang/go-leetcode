package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}))
}

// leetcode1436_旅行终点站
func destCity(paths [][]string) string {
	m := make(map[string]int)
	for i := 0; i < len(paths); i++ {
		m[paths[i][1]] -= 1
		m[paths[i][0]] += 1

	}
	for key, value := range m {
		if value == -1 {
			return key
		}
	}
	return ""
}
