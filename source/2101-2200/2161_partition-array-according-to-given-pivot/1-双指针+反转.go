package main

func main() {

}

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = pivot
	}
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			res[left] = nums[i]
			left++
		} else if nums[i] > pivot {
			res[right] = nums[i]
			right--
		}
	}
	reverse(res[right+1:]) // 反转后面顺序
	return res
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
