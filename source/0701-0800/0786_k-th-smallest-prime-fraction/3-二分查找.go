package main

func main() {

}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0
	for {
		mid := left + (right-left)/2
		count := 0
		x, y := 0, 1 // 记录最大的分子/分母
		i := -1
		for j := 1; j < n; j++ {
			for float64(arr[i+1])/float64(arr[j]) < mid { // 小于目标
				i++
				if arr[i]*y > arr[j]*x { // 更新：a/b > c/d => a*d > bc 保存a,b
					x, y = arr[i], arr[j]
				}
			}
			count = count + (i + 1) // 除以当前arr[j],总计几个数小于mid
		}
		if count == k {
			return []int{x, y}
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}
}
