package main

func main() {

}

// leetcode2201_统计可以提取的工件
func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	res := 0
	m := make(map[[2]int]int)
	for i := 0; i < len(dig); i++ {
		a, b := dig[i][0], dig[i][1]
		m[[2]int{a, b}] = 1
	}
	for i := 0; i < len(artifacts); i++ {
		a, b, c, d := artifacts[i][0], artifacts[i][1], artifacts[i][2], artifacts[i][3]
		flag := true
		for x := a; x <= c; x++ {
			for y := b; y <= d; y++ {
				if m[[2]int{x, y}] == 0 {
					flag = false
				}
			}
		}
		if flag == true {
			res++
		}
	}
	return res
}
