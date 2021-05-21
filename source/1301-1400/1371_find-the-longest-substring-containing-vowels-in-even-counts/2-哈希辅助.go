package main

func main() {

}

// leetcode1371_每个元音包含偶数次的最长子字符串
func findTheLongestSubstring(s string) int {
	res := 0
	m := make(map[[5]int]int)
	status := [5]int{}
	m[status] = -1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status[0] = 1 - status[0]
		case 'e':
			status[1] = 1 - status[1]
		case 'i':
			status[2] = 1 - status[2]
		case 'o':
			status[3] = 1 - status[3]
		case 'u':
			status[4] = 1 - status[4]
		}
		if v, ok := m[status]; ok {
			res = max(res, i-v)
		} else {
			m[status] = i // 记录第一次出现的位置
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
