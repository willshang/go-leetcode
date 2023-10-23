package main

func main() {

}

// leetcode2120_执行所有后缀指令
func executeInstructions(n int, startPos []int, s string) []int {
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		count := 0
		a, b := startPos[0], startPos[1]
		for j := i; j < len(s); j++ {
			switch s[j] {
			case 'L':
				b--
			case 'R':
				b++
			case 'U':
				a--
			case 'D':
				a++
			}
			if 0 <= a && a < n && 0 <= b && b < n {
				count++
			} else {
				break
			}
		}
		res = append(res, count)
	}
	return res
}
