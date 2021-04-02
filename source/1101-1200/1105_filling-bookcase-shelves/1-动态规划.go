package main

func main() {

}

// leetcode1105_填充书架
func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)
	dp := make([]int, n+1) // 以第i本书作为结尾的总高度
	for i := 1; i <= n; i++ {
		w, h := books[i-1][0], books[i-1][1]
		dp[i] = dp[i-1] + h // 当前这本书单独一层的高度
		for j := i - 1; j > 0; j-- {
			if w+books[j-1][0] <= shelf_width {
				w = w + books[j-1][0]
				h = max(h, books[j-1][1])
				dp[i] = min(dp[i], dp[j-1]+h)
			} else {
				break
			}
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
