package main

func main() {

}

// leetcode1897_重新分配字符使所有字符串都相等
func makeEqual(words []string) bool {
	arr := [26]int{}
	n := len(words)
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			arr[int(words[i][j]-'a')]++
		}
	}
	for i := 0; i < 26; i++ {
		if arr[i]%n != 0 {
			return false
		}
	}
	return true
}
