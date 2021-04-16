package main

func main() {

}

func findTheWinner(n int, k int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	last := 0
	for len(arr) > 1 {
		index := (last + k - 1) % len(arr)
		arr = remove(arr, index)
		last = index
	}
	return arr[0] + 1
}

func remove(arr []int, index int) []int {
	if index == 0 {
		return arr[1:]
	}
	if index == len(arr)-1 {
		return arr[:len(arr)-1]
	}
	return append(arr[:index], arr[index+1:]...)
}
