package main

func main() {

}

// leetcode1255_得分最高的单词集合
var res int
var count [26]int

func maxScoreWords(words []string, letters []byte, score []int) int {
	res = 0
	count = [26]int{}
	for i := 0; i < len(letters); i++ {
		count[int(letters[i]-'a')]++ // 统计字符出现的次数
	}
	dfs(words, score, 0, [26]int{})
	return res
}

func dfs(words []string, score []int, index int, arr [26]int) {
	sum := 0
	for i := 0; i < 26; i++ {
		if arr[i] > count[i] {
			return
		}
		sum = sum + arr[i]*score[i]
	}
	res = max(res, sum)
	for i := index; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			arr[int(words[i][j]-'a')]++
		}
		dfs(words, score, i+1, arr)
		for j := 0; j < len(words[i]); j++ {
			arr[int(words[i][j]-'a')]--
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
