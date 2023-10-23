package main

func main() {

}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)/2   // 枚举左边
		if x-arr[mid] > arr[mid+k]-x { // 看x离left和right哪个远
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
