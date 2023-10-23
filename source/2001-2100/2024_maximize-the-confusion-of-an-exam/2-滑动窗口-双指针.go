package main

func main() {

}

// leetcode2024_考试的最大困扰度
func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	res := 0
	left, right := 0, 0
	countA, countB := 0, 0
	for ; right < n; right++ {
		if answerKey[right] == 'T' {
			countA++
		} else {
			countB++
		}
		for countA > k && countB > k { // 都大于k时，滑动窗口才移动
			if answerKey[left] == 'T' {
				countA--
			} else {
				countB--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
