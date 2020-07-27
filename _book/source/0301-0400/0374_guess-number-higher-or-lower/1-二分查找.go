package main

func main() {

}

func guess(num int) int {
	return 0
}

// leetcode374_猜数字大小
func guessNumber(n int) int {
	low := 1
	high := n
	for low < high {
		mid := low + (high-low)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
