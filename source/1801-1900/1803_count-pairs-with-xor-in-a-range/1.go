package main

func main() {

}

func countPairs(nums []int, low int, high int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {

	}

	return res
}

type Trie struct {
	next   [2]*Trie // 的
	ending int      // 次数
}

// 插入word
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [2]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

// 查找
func (this *Trie) Search(word string) bool {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}
