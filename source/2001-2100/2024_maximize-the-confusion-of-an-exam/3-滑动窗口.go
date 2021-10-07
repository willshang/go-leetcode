package main

func main() {

}

func maxConsecutiveAnswers(answerKey string, k int) int {
	arr := []byte(answerKey)
	return max(longestOnes(arr, k, 'T'), longestOnes(arr, k, 'F'))
}

func longestOnes(A []byte, K int, target byte) int {
	res := 0
	left, right := 0, 0
	count := 0
	for right = 0; right < len(A); right++ {
		if A[right] == target {
			count++
		}
		for count > K {
			if A[left] == target {
				count--
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
