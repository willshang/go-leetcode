package main

import "fmt"

func main() {
	fmt.Println(getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
}

// leetcode1487_保证文件名唯一
func getFolderNames(names []string) []string {
	m := make(map[string]int)
	for i, name := range names {
		if value, ok := m[name]; ok {
			names[i] = getName(m, name, value)
			m[names[i]] = 1
		} else {
			m[name] = 1
		}
	}
	return names
}

func getName(m map[string]int, name string, n int) string {
	newName := name + fmt.Sprintf("(%d)", n)
	if _, ok := m[newName]; ok {
		return getName(m, name, n+1)
	}
	m[name] = n + 1
	return newName
}
