package main

func main() {

}

// leetcode1400_构造K个回文字符串
func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			res++
		}
	}
	return res <= k
}
