package main

func main() {

}

func minimumPerimeter(neededApples int64) int64 {
	var res, left, right int64
	left = 1
	right = 100000
	for left <= right {
		mid := left + (right-left)/2
		total := 2 * mid * (mid + 1) * (2*mid + 1)
		if total >= neededApples {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res * 8
}
