package main

func main() {

}

func maxRepOpt1(text string) int {
	n := len(text)
	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		m[text[i]]++ // 统计每个字母出现的次数
	}

}
