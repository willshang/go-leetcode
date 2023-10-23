package main

func main() {

}

type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{m: map[int][]string{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := 0; i < len(dictionary); i++ {
		this.m[len(dictionary[i])] = append(this.m[len(dictionary[i])], dictionary[i])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if len(this.m[len(searchWord)]) == 0 {
		return false
	}
	for i := 0; i < len(this.m[len(searchWord)]); i++ {
		word := this.m[len(searchWord)][i]
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
