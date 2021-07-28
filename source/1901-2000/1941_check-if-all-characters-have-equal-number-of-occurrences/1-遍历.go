package main

func main() {

}

func areOccurrencesEqual(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	count := 0
	for _, v := range m {
		if count == 0 {
			count = v
		} else {
			if count != v {
				return false
			}
		}
	}
	return true
}
