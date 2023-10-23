package main

func main() {

}

// leetcode916_单词子集
func wordSubsets(A []string, B []string) []string {
	maxArr := make([]int, 26) // B中的字母的最大频次
	for i := 0; i < len(B); i++ {
		temp := count(B[i])
		for i := 0; i < 26; i++ {
			maxArr[i] = max(maxArr[i], temp[i])
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(A); i++ {
		temp := count(A[i])
		flag := true
		for i := 0; i < 26; i++ {
			if temp[i] < maxArr[i] {
				flag = false
				break
			}
		}
		if flag == true {
			res = append(res, A[i])
		}
	}
	return res
}

func count(str string) [26]int {
	arr := [26]int{}
	for i := 0; i < len(str); i++ {
		arr[str[i]-'a']++
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
