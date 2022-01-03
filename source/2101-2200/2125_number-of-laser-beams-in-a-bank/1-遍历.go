package main

func main() {

}

func numberOfBeams(bank []string) int {
	res := 0
	prev := 0
	n, m := len(bank), len(bank[0])
	for i := 0; i < n; i++ {
		cur := 0
		for j := 0; j < m; j++ {
			if bank[i][j] == '1' {
				cur++
			}
		}
		if cur == 0 {
			continue
		}
		res = res + cur*prev
		prev = cur
	}
	return res
}
