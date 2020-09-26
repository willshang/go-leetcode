package main

func main() {

}

// leetcode211_添加与搜索单词-数据结构设计
type Trie struct {
	next   [26]*Trie
	ending int
}

func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [26]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

func (this *Trie) Search(word string, k int) bool {
	temp := this
	for i := k; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < len(temp.next); j++ {
				if temp.next[j] != nil && temp.next[j].Search(word, i+1) {
					return true
				}
			}
			return false
		}
		value := word[i] - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{trie: &Trie{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word, 0)
}
