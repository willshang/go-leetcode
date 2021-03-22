package main

import "fmt"

func main() {
	fmt.Println(simplifiedFractions(3))
}

func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	m := make(map[string]bool)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			str := fmt.Sprintf("%d/%d", j, i)
			if _, ok := m[str]; ok {
				continue
			}
			res = append(res, str)
			for k := 1; i*k <= n; k++ {
				m[fmt.Sprintf("%d/%d", j*k, i*k)] = true
			}
		}
	}
	return res
}
