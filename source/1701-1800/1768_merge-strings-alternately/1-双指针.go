package main

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		res = res + string(word1[i])
		res = res + string(word2[j])
		i++
		j++
	}
	if i < len(word1) {
		res = res + string(word1[i:])
	}
	if j < len(word2) {
		res = res + string(word2[j:])
	}
	return res
}
