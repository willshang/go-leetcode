package main

func main() {

}

func maxDistance(colors []int) int {
	n := len(colors)
	if colors[0] != colors[n-1] {
		return n - 1
	}
	left, right := 1, n-2
	for colors[n-1] == colors[left] {
		left++
	}
	for colors[0] == colors[right] {
		right--
	}
	return max(right, n-1-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
