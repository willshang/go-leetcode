package main

func main() {

}

type Trie struct {
	next   map[byte]*Trie
	ending int
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next:   make(map[byte]*Trie),
		ending: 0,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := byte(v - 'a')
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   make(map[byte]*Trie),
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	temp := this
	for _, v := range word {
		value := byte(v - 'a')
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for _, v := range prefix {
		value := byte(v - 'a')
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	return true
}
