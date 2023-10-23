package main

func main() {

}

// leetcode1542_找出最长的超赞子字符串
func longestAwesome(s string) int {
	res := 0
	n := len(s)
	m := make(map[int]int) // 保存每个状态第1次出现的下标
	m[0] = -1              // 0对应的下标
	cur := 0
	for i := 0; i < n; i++ {
		value := int(s[i] - '0')
		cur = cur ^ (1 << value)
		if index, ok := m[cur]; ok { // 相同的情况
			res = max(res, i-index)
		} else {
			m[cur] = i
		}
		for j := 0; j < 10; j++ { // 相差1位的情况
			if index, ok := m[cur^(1<<j)]; ok {
				res = max(res, i-index)
			}
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
