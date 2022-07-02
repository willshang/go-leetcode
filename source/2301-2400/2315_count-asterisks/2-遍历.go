package main

func main() {

}

func countAsterisks(s string) int {
	res := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			count++
		}
		if s[i] == '*' && count%2 == 0 {
			res++
		}
	}
	return res
}
