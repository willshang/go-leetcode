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
	for i := 0; i < len(paths); i++ {
		flag := false
		for j := 0; j < len(paths); j++ {
			if j == i {
				continue
			}
			if paths[i][1] == paths[j][0] {
				flag = true
				break
			}
		}
		if flag == false {
			return paths[i][1]
		}
	}
	return ""
}
