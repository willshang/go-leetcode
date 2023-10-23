package main

func main() {

}

// leetcode466_统计重复个数
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	count := 0
	j := 0
	for a := 0; a < n1; a++ {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				if j == len(s2)-1 {
					j = 0
					count++
				} else {
					j++
				}
			}
		}
	}
	return count / n2
}
