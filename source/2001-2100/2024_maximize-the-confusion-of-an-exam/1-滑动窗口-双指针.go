package main

func main() {

}

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if answerKey[i] == 'T' {
			arr[i] = 1
		}
	}
	a := longestOnes(arr, k)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		if answerKey[i] == 'F' {
			arr[i] = 1
		}
	}
	b := longestOnes(arr, k)
	return max(a, b)
}

// leetcode 1004.最大连续1的个数III
func longestOnes(A []int, K int) int {
	res := 0
	left, right := 0, 0
	count := 0
	for right = 0; right < len(A); right++ {
		if A[right] == 0 {
			count++
		}
		for count > K {
			if A[left] == 0 {
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
