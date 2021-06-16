package main

func main() {

}

func isCovered(ranges [][]int, left int, right int) bool {
	arr := make([]int, 52)
	for i := 0; i < len(ranges); i++ {
		a, b := ranges[i][0], ranges[i][1]
		arr[a]++
		arr[b+1]--
	}
	sum := 0
	for i := 0; i <= 50; i++ {
		sum = sum + arr[i]
		if left <= i && i <= right && sum <= 0 {
			return false
		}
	}
	return true
}
