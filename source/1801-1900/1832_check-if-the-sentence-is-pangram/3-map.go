package main

func main() {

}

func checkIfPangram(sentence string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(sentence); i++ {
		m[sentence[i]-'a']++
	}
	return len(m) == 26
}
