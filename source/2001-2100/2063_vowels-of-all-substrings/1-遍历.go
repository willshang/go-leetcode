package main

func main() {

}

// leetcode2063_所有子字符串中的元音
var m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func countVowels(word string) int64 {
	n := len(word)
	res := int64(0)
	for i := 0; i < n; i++ {
		if m[word[i]] == true {
			// 计算当前字符出现在多少种组合里面
			// 总数：left * right
			res = res + int64(i+1)*int64(n-i) // 左边个数（包含当前）*右边个数
		}
	}
	return res
}
