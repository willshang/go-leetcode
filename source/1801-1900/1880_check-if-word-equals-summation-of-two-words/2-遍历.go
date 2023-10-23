package main

func main() {

}

// leetcode1880_检查某单词是否等于两单词之和
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return getValue(firstWord)+getValue(secondWord) == getValue(targetWord)
}

func getValue(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		res = res*10 + int(str[i]-'a')
	}
	return res
}
