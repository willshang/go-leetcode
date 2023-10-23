package main

func main() {

}

func prefixCount(words []string, pref string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= len(pref) && words[i][:len(pref)] == pref {
			res++
		}
	}
	return res
}
