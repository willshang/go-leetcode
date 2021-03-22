package main

func main() {

}

func minOperations(s string) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0')%2 != i%2 {
			a++
		} else {
			b++
		}
	}
	if a > b {
		return b
	}
	return a
}
