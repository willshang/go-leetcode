package main

func main() {

}

// leetcode927_三等分
func threeEqualParts(arr []int) []int {
	res := []int{-1, -1}
	n := len(arr)
	indexArr := make([]int, 0)
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			indexArr = append(indexArr, i)
		}
	}
	count := len(indexArr)
	if count == 0 {
		return []int{0, 2}
	}
	if count%3 != 0 {
		return res
	}
	// 3个部分，每个部分的1的个数都是一样的
	a, b := count/3, count/3*2 // 第2、3组开始的位置
	i, j, k := indexArr[0], indexArr[a], indexArr[b]
	for k < n && arr[i] == arr[j] && arr[j] == arr[k] {
		i++
		j++
		k++
	}
	if k == n {
		return []int{i - 1, j}
	}
	return res
}
