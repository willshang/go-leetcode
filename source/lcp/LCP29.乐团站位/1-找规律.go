package main

func main() {

}

// leetcode_lcp29_乐团站位
func orchestraLayout(num int, xPos int, yPos int) int {
	x := min(xPos, num-1-xPos)
	y := min(yPos, num-1-yPos)
	k := min(x, y) // 在第几圈（从0开始）
	// n*n - (n-2k)*(n-2k) = n*n-(n*n+4k*k-4nk)= 4nk-4k*k
	total := 4 * k * (num - k) % 9 // 第几圈外总共的个数
	if xPos == k {                 // 上边
		return (total+yPos-k)%9 + 1
	} else if yPos == num-1-k { // 右边
		before := num - 2*k - 1
		return (total+before+xPos-k)%9 + 1
	} else if xPos == num-1-k { // 下边
		before := (num - 2*k - 1) * 2
		return (total+before+num-k-1-yPos)%9 + 1
	} else if yPos == k { // 左边
		before := (num - 2*k - 1) * 3
		return (total+before+num-k-1-xPos)%9 + 1
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
