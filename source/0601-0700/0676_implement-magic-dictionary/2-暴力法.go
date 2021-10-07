package main

func main() {

}

// leetcode676_实现一个魔法字典
type MagicDictionary struct {
	arr []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{arr: make([]string, 0)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.arr = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for i := 0; i < len(this.arr); i++ {
		word := this.arr[i]
		if len(word) != len(searchWord) {
			continue
		}
		count := 0
		for j := 0; j < len(searchWord); j++ {
			if word[j] != searchWord[j] {
				count++
				if count > 1 {
					break
				}
			}
		}
		if count == 1 {
			return true
		}
	}
	return false
}
