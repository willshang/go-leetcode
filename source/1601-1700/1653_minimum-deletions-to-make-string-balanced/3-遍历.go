package main

func main() {

}

func minimumDeletions(s string) int {
	aCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount = aCount + 1
		}
	}
	res := aCount
	bCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount--
		} else {
			bCount++
		}
		// 要删除：前面b的个数+后面a的个数
		res = min(res, aCount+bCount)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
