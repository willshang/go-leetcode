package main

func main() {

}

func countSubstrings(s string, t string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			diff := 0
			for k := 0; i+k < len(s) && j+k < len(t); k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				}
				if diff == 1 {
					res++
				}
			}
		}
	}
	return res
}
