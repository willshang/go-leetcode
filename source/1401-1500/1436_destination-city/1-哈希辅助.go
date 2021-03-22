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
	m := make(map[string]string)
	for i := 0; i < len(paths); i++ {
		m[paths[i][0]] = paths[i][1]
	}
	for _, v := range m {
		if _, ok := m[v]; !ok {
			return v
		}
	}
	return ""
}
