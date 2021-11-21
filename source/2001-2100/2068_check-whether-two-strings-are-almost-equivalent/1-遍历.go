package main

func main() {

}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	a, b := [26]int{}, [26]int{}
	for i := 0; i < len(word1); i++ {
		a[int(word1[i]-'a')]++
	}
	for i := 0; i < len(word2); i++ {
		b[int(word2[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if a[i]-b[i] > 3 || b[i]-a[i] > 3 {
			return false
		}
	}
	return true
}
