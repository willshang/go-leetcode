package main

import (
	"strings"
)

func main() {

}

func removeSubfolders(folder []string) []string {
	res := make([]string, 0)
	m := make(map[string]bool)
	for i := 0; i < len(folder); i++ {
		m[folder[i]] = true
	}
	arr := make([]int, len(folder))
	for i := 0; i < len(folder); i++ {
		for j := 0; j < len(folder[i]); j++ {
			if folder[i][j] == '/' && j > 0 && m[folder[i][:j]] == true {
				arr[i] = 1
				break
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			res = append(res, folder[i])
		}
	}
	return res
}
