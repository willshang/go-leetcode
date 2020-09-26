package main

func main() {

}

type WordsFrequency struct {
	ending int
	next   [26]*WordsFrequency
}

func Constructor(book []string) WordsFrequency {
	res := WordsFrequency{}
	for _, v := range book {
		res.Insert(v)
	}
	return res
}

func (this *WordsFrequency) Get(word string) int {
	temp := this
	for _, v := range word {
		nextWord := v - 'a'
		if temp.next[nextWord] == nil {
			return 0
		}
		temp = temp.next[nextWord]
	}
	return temp.ending
}

func (this *WordsFrequency) Insert(word string) {
	temp := this
	for _, v := range word {
		nextWord := v - 'a'
		if temp.next[nextWord] == nil {
			temp.next[nextWord] = &WordsFrequency{}
		}
		temp = temp.next[nextWord]
	}
	temp.ending = temp.ending + 1
}
