package main

func main() {

}

func checkIfPangram(sentence string) bool {
	res := 0
	target := 1<<26 - 1
	for i := 0; i < len(sentence); i++ {
		res = res | 1<<(sentence[i]-'a')
	}
	return res == target
}
