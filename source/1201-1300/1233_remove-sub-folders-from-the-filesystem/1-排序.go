package main

import (
	"sort"
	"strings"
)

func main() {

}

// leetcode1233_删除子文件夹
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := make([]string, 0)
	res = append(res, folder[0])
	prev := folder[0]
	for i := 1; i < len(folder); i++ {
		// 加'/'避免/a/b => /a/bb的情况
		if strings.HasPrefix(folder[i], prev+"/") == false {
			res = append(res, folder[i])
			prev = folder[i]
		}
	}
	return res
}
