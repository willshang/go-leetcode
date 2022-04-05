package main

func main() {

}

// leetcode2186_使两字符串互为字母异位词的最少步骤数
func minSteps(s string, t string) int {
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[int(s[i]-'a')]++
	}
	for i := 0; i < len(t); i++ {
		arr[int(t[i]-'a')]--
	}
	res := 0
	for i := 0; i < 26; i++ {
		res = res + abs(arr[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
