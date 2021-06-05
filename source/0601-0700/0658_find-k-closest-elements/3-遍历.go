package main

func main() {

}

// leetcode658_找到K个最接近的元素
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for i := 1; i <= len(arr)-k; i++ {
		if x-arr[left] <= arr[right]-x { // 看x离left和right哪个远，远的就移动 相等就right-1
			right--
		} else {
			left++
		}
	}
	return arr[left : right+1]
}
