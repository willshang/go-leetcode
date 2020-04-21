package main

func main() {

}

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) == false {
			low = mid + 1
		} else if isBadVersion(mid) == true && isBadVersion(mid-1) == true {
			high = mid - 1
		} else if isBadVersion(mid) == true && isBadVersion(mid-1) == false {
			return mid
		}
	}
	return -1
}
