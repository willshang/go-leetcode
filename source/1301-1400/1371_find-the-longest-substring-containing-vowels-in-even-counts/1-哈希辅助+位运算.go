package main

func main() {

}

func findTheLongestSubstring(s string) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	status := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status = status ^ 0b00001
		case 'e':
			status = status ^ 0b00010
		case 'i':
			status = status ^ 0b00100
		case 'o':
			status = status ^ 0b01000
		case 'u':
			status = status ^ 0b10000
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
