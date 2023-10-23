package main

func main() {

}

func customSortString(S string, T string) string {
	count := make([]int, 26)
	for i := 0; i < len(T); i++ {
		count[T[i]-'a']++
	}
	res := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		for j := 0; j < count[S[i]-'a']; j++ {
			res = append(res, S[i])
		}
		count[S[i]-'a'] = 0
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < count[i]; j++ {
			res = append(res, byte(i+'a'))
		}
	}
	return string(res)
}
