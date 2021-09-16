package main

func main() {

}

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
	if index >= len(words) {
		return
	}
	dfs(words, score, index+1, arr) // 不选
	for j := 0; j < len(words[index]); j++ {
		arr[int(words[index][j]-'a')]++
	}
	dfs(words, score, index+1, arr) // 选
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
