package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}))
}

func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for i := 0; i < len(paths); i++ {
		m[paths[i][1]] = true
	}
	for i := 0; i < len(paths); i++ {
		m[paths[i][0]] = false
	}
	for key, value := range m {
		if value == true {
			return key
		}
	}
	return ""
}
