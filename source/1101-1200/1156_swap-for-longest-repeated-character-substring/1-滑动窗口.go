package main

func main() {

}

// leetcode1156_单字符重复子串的最大长度
func maxRepOpt1(text string) int {
	res := 0
	n := len(text)
	arr := [26]int{}
	for i := 0; i < n; i++ {
		v := int(text[i] - 'a')
		arr[v]++ // 统计每个字母出现的次数
	}
	for i := 0; i < n; {
		v := int(text[i] - 'a')
		countA := 0 // 统计相同的个数
		for i+countA < n && text[i] == text[i+countA] {
			countA++
		}
		j := i + countA + 1
		countB := 0 // 统计相隔1个不同字符往后连续的次数
		for j+countB < n && text[i] == text[j+countB] {
			countB++
		}
		// 2种情况
		// 1、a...axa...ay..a 可以拿1个来补齐:+1
		// 2、没有多余的补齐，可以拿第二段右侧最后一个点或者第一段左侧第一个点来补:+0
		total := min(countA+countB+1, arr[v]) // 取较少的
		res = max(res, total)                 // 更新结果
		i = i + countA                        // 窗口后移

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
