package main

func main() {

}

func maximumSubsequenceCount(text string, pattern string) int64 {
	a, b := pattern[0], pattern[1]
	n := len(text)
	var countA, countB int64
	res := int64(0)
	for i := 0; i < n; i++ {
		if text[i] == b { // 存在相同的情况，先判断b
			countB++
			res = res + countA // 子序列数量：加上前面的a
		}
		if text[i] == a {
			countA++
		}
	}
	return res + max(countA, countB) // 最后决定插入a还是b：插入a在最前面，插入b在最后面
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
