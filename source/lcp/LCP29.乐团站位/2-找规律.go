package main

func main() {

}

func orchestraLayout(num int, xPos int, yPos int) int {
	x := min(xPos, num-1-xPos)
	y := min(yPos, num-1-yPos)
	k := min(x, y) // 在第几圈（从0开始）
	// n*n - (n-2k)*(n-2k) = n*n-(n*n+4k*k-4nk)= 4nk-4k*k
	if xPos <= yPos {
		total := num*num - (num-2*k)*(num-2*k)
		return (total+xPos-k+yPos-k)%9 + 1
	} else {
		total := num*num - (num-(2*k+2))*(num-(2*k+2))
		return (total-(xPos-k)-(yPos-k))%9 + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
