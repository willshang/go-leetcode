package main

func main() {

}

// leetcode1961_检查字符串是否为数组前缀
func isPrefixString(s string, words []string) bool {
	temp := ""
	for i := 0; i < len(words); i++ {
		temp = temp + words[i]
		if temp == s {
			return true
		}
	}
	return false
}
