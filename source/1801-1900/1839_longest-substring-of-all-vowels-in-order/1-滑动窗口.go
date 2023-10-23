package main

func main() {

}

// leetcode1839_所有元音按顺序排布的最长子字符串
func longestBeautifulSubstring(word string) int {
	res := 0
	count := 1
	i := 0
	for j := 1; j < len(word); j++ { // 题目保证了全是目标元素，而且顺序是有序的
		if word[j] < word[j-1] { // 大小不对，窗口右移，重新计数
			i = j
			count = 1
		} else if word[j] > word[j-1] {
			count++
		}
		if count == 5 { // 长度为5
			res = max(res, j-i+1)
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
