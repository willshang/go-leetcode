package main

func main() {

}

// leetcode1828_统计一个圆中点的数目
func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < len(points); j++ {
			if judge(queries[i], points[j]) == true {
				count++
			}
		}
		res[i] = count
	}
	return res
}

func judge(query []int, point []int) bool {
	x, y, r := query[0], query[1], query[2]
	x1, y1 := point[0], point[1]
	return (x-x1)*(x-x1)+(y-y1)*(y-y1) <= r*r
}
