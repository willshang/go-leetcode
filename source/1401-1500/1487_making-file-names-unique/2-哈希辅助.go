package main

import "fmt"

func main() {
	fmt.Println(getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
}

// leetcode1487_go_保证文件名唯一
func getFolderNames(names []string) []string {
	m := make(map[string]int)
	res := make([]string, 0)
	for _, name := range names {
		if value, ok := m[name]; ok {
			for {
				new := name + fmt.Sprintf("(%d)", value)
				if _, ok2 := m[new]; ok2 {
					value++
					continue
				}
				res = append(res, new)
				m[new] = 1
				m[name] = value
				break
			}
		} else {
			res = append(res, name)
			m[name] = 1
		}
	}
	return res
}
