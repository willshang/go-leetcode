package main

func main() {

}

// leetcode1943_描述绘画结果
func splitPainting(segments [][]int) [][]int64 {
	res := make([][]int64, 0)
	arr := make([]int64, 100005) // 和
	m := make([]bool, 100005)
	n := len(segments)
	for i := 0; i < n; i++ {
		start := segments[i][0]
		end := segments[i][1]
		count := segments[i][2]
		arr[start] = arr[start] + int64(count)
		arr[end] = arr[end] - int64(count)
		m[start] = true
		m[end] = true
	}
	var sum int64
	var prev int64
	for i := 1; i < 100005; i++ {
		if m[i] == true {
			if sum > 0 {
				res = append(res, []int64{prev, int64(i), sum})
			}
			sum = sum + arr[i]
			prev = int64(i)
		}
	}
	return res
}
