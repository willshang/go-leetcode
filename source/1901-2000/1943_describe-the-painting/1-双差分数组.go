package main

func main() {

}

func splitPainting(segments [][]int) [][]int64 {
	res := make([][]int64, 0)
	arr1 := make([]int64, 100005) // 和
	arr2 := make([]int64, 100005) // 次数
	n := len(segments)
	for i := 0; i < n; i++ {
		start := segments[i][0]
		end := segments[i][1]
		count := segments[i][2]
		arr1[start], arr2[start] = arr1[start]+int64(count), arr1[start]+1
		arr1[end], arr2[end] = arr1[end]-int64(count), arr1[end]-1
	}
	prevIndex := -1
	var prev1, prev2 int64
	var sum1, sum2 int64
	for i := 1; i < 100005; i++ {
		sum1 = sum1 + arr1[i]
		sum2 = sum2 + arr2[i]
		if sum1 > 0 {
			if prevIndex == -1 { // 之前不为-1，开头
				prevIndex = i
				prev1, prev2 = sum1, sum2
			} else if prev1 != sum1 { // 和不同，加入区间
				res = append(res, []int64{int64(prevIndex), int64(i), prev1})
				prevIndex = i
				prev1, prev2 = sum1, sum2
			} else {
				if prev2 != sum2 { // 次数不同。加入区间
					res = append(res, []int64{int64(prevIndex), int64(i), prev1})
					prevIndex = i
					prev1, prev2 = sum1, sum2
				}
			}
		} else if sum1 == 0 {
			if prevIndex > 0 { // 和为0，之前不为0，加入区间
				res = append(res, []int64{int64(prevIndex), int64(i), prev1})
				prevIndex = -1
				prev1, prev2 = sum1, sum2
			}
		}
	}
	return res
}
