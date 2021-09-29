package main

func main() {

}

// leetcode_lcp46_志愿者调配
type Node struct {
	a float64 // 未知final[0] x的系数
	b int     // 已知的系数
}

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	arr := make([][]int, n) // 邻接表
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	final := make([]Node, 0)
	final = append(final, Node{a: 1, b: 0}) // final为最终的结果，a为final[0]未知x的系数
	for i := 0; i < len(finalCnt); i++ {
		final = append(final, Node{a: 0, b: finalCnt[i]})
	}
	for i := len(plans) - 1; i >= 0; i-- { // 倒推
		a, b := plans[i][0], plans[i][1]
		if a == 1 {
			final[b] = Node{
				a: final[b].a * 2,
				b: final[b].b * 2,
			}
		} else if a == 2 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				final[next] = Node{
					a: final[next].a - final[b].a,
					b: final[next].b - final[b].b,
				}
			}
		} else if a == 3 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				final[next] = Node{
					a: final[next].a + final[b].a,
					b: final[next].b + final[b].b,
				}
			}
		}
	}
	a, b := float64(0), int64(0)
	for i := 0; i < len(final); i++ {
		a = a + final[i].a
		b = b + int64(final[i].b)
	}
	per := float64(totalNum-b) / a // 求x的值
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(final[i].a*per) + final[i].b
	}
	return res
}
