package main

func main() {

}

// leetcode1832_判断句子是否为全字母句
func checkIfPangram(sentence string) bool {
	arr := [26]int{}
	for i := 0; i < len(sentence); i++ {
		arr[int(sentence[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if arr[i] == 0 {
			return false
		}
	}
	return true
}
