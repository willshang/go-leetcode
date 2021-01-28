package main

import "strings"

func main() {

}

// leetcode609_在系统中查找重复文件
func findDuplicate(paths []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)
	for i := 0; i < len(paths); i++ {
		arr := strings.Split(paths[i], " ")
		for j := 1; j < len(arr); j++ {
			index := strings.LastIndexByte(arr[j], '(')
			content := arr[j][index+1 : len(arr[j])-1]
			m[content] = append(m[content], arr[0]+"/"+arr[j][:index])
		}
	}
	for _, v := range m {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
