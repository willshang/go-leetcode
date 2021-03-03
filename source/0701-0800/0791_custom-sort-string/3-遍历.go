package main

func main() {

}

func customSortString(S string, T string) string {
	res := []byte(T)
	index := 0
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(res); j++ {
			if res[j] == S[i] {
				res[j], res[index] = res[index], res[j]
				index++
			}
		}
	}
	return string(res)
}
