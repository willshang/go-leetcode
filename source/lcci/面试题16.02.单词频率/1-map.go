package main

func main() {

}

// 程序员面试金典16.02_单词频率
type WordsFrequency struct {
	m map[string]int
}

func Constructor(book []string) WordsFrequency {
	res := WordsFrequency{m: make(map[string]int)}
	for k := range book {
		res.m[book[k]]++
	}
	return res
}

func (this *WordsFrequency) Get(word string) int {
	return this.m[word]
}
