package main

func main() {

}

func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		if boxes[i] == '1' {
			arr = append(arr, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(arr); j++ {
			res[i] = res[i] + abs(arr[j]-i)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
