package main

func main() {

}

func maxScoreWords(words []string, letters []byte, score []int) int {
	count := [26]int{}
	for i := 0; i < len(letters); i++ {
		count[int(letters[i]-'a')]++ // 统计字符出现的次数
	}
	n := len(words)
	total := 1 << n // 总状态数
	res := 0
	for i := 0; i < total; i++ { // 枚举所有状态
		arr := getStatus(words, i)                  // 统计每种状态所需字符个数
		res = max(res, getScore(score, count, arr)) // 计算得分
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 统计该状态每个字符多少个
func getStatus(words []string, status int) [26]int {
	arr := [26]int{}
	for i := 0; i < len(words); i++ {
		if status&(1<<i) > 0 {
			for j := 0; j < len(words[i]); j++ {
				arr[int(words[i][j]-'a')]++
			}
		}
	}
	return arr
}

func getScore(score []int, count, arr [26]int) int {
	res := 0
	for i := 0; i < 26; i++ {
		if count[i] < arr[i] {
			return 0
		}
		res = res + score[i]*arr[i]
	}
	return res
}
