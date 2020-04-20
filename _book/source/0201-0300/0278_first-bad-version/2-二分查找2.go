package main

func main() {

}

func isBadVersion(version int) bool {
	return true
}

// leetcode278_第一个错误的版本
func firstBadVersion(n int) int {
	low := 1
	high := n
	for low < high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
