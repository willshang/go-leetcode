package main

func main() {

}

func canBeIncreasing(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if judge(nums, i) == true {
			return true
		}
	}
	return false
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
