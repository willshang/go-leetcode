package main

func main() {

}

// leetcode1909_删除一个元素使数组严格递增
func canBeIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return judge(nums, i) == true || judge(nums, i+1) == true
		}
	}
	return true
}

func judge(nums []int, index int) bool {
	arr := append([]int{}, nums[:index]...)
	arr = append(arr, nums[index+1:]...)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
