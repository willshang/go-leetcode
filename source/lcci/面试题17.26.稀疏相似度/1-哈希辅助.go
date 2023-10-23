package main

import "fmt"

func main() {

}

// 程序员面试金典17.26_稀疏相似度
func computeSimilarities(docs [][]int) []string {
	res := make([]string, 0)
	n := len(docs)
	m := make(map[[2]int]int)
	m1 := make(map[int][]int) // 字符出现的位置
	for i := 0; i < n; i++ {
		for j := 0; j < len(docs[i]); j++ {
			char := docs[i][j]
			for _, v := range m1[char] {
				m[[2]int{v, i}]++
			}
			m1[char] = append(m1[char], i)
		}
	}
	for k, v := range m {
		x := v
		y := len(docs[k[0]]) + len(docs[k[1]]) - v
		res = append(res, fmt.Sprintf("%d,%d: %.4f",
			k[0], k[1], float64(x)/float64(y)+1e-9))
	}
	return res
}
