package main

func main() {

}

// leetcode835_图像重叠
func largestOverlap(img1 [][]int, img2 [][]int) int {
	res := 0
	n := len(img1)
	arr := make([][]int, 2*n+1)
	for i := 0; i <= 2*n; i++ {
		arr[i] = make([]int, 2*n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] == 1 {
				for a := 0; a < n; a++ {
					for b := 0; b < n; b++ {
						if img2[a][b] == 1 {
							arr[i-a+n][j-b+n]++ // i-a，j-b是移动的，+n修正负数
						}
					}
				}
			}
		}
	}
	for i := 0; i <= 2*n; i++ {
		for j := 0; j <= 2*n; j++ {
			res = max(res, arr[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
