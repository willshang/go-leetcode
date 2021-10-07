package main

func main() {

}

type MagicDictionary struct {
	next   [26]*MagicDictionary // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int                  // 次数（可以改为bool）
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		next:   [26]*MagicDictionary{},
		ending: 0,
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := 0; i < len(dictionary); i++ {
		word := dictionary[i]
		temp := this
		for _, v := range word {
			value := v - 'a'
			if temp.next[value] == nil {
				temp.next[value] = &MagicDictionary{
					next:   [26]*MagicDictionary{},
					ending: 0,
				}
			}
			temp = temp.next[value]
		}
		temp.ending++
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	cur := this
	arr := []byte(searchWord)
	for i := 0; i < len(searchWord); i++ {
		b := searchWord[i]
		for j := 0; j < 26; j++ {
			if j+'a' == int(b) {
				continue
			}
			arr[i] = byte('a' + j)
			if cur.SearchWord(string(arr[i:])) == true {
				return true
			}
		}
		arr[i] = b
		if cur.next[int(b-'a')] == nil {
			return false
		}
		cur = cur.next[int(b-'a')]
	}
	return false
}

func (this *MagicDictionary) SearchWord(word string) bool {
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
