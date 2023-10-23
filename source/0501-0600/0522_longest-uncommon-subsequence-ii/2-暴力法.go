package main

func main() {

}

// leetcode522_最长特殊序列II
func findLUSlength(strs []string) int {
	res := -1
	var j int
	for i := 0; i < len(strs); i++ {
		for j = 0; j < len(strs); j++ {
			if i == j {
				continue
			}
			if judge(strs[i], strs[j]) == true {
				break
			}
		}
		if j == len(strs) && len(strs[i]) > res {
			res = len(strs[i])
		}
	}

	return res
}

func judge(a, b string) bool {
	j := 0
	for i := 0; i < len(b) && j < len(a); i++ {
		if a[j] == b[i] {
			j++
		}
	}
	return j == len(a)
}
