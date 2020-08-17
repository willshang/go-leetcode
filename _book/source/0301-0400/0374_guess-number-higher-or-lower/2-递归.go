package main

func main() {

}

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	return binary(1, n)
}

func binary(left, right int) int {
	mid := left + (right-left)/2
	if guess(mid) == 1 {
		return binary(mid+1, right)
	} else if guess(mid) == -1 {
		return binary(left, mid-1)
	} else {
		return mid
	}
}
