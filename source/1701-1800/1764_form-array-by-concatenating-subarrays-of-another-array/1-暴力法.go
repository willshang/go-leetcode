package main

func main() {

}

// leetcode1764_通过连接另一个数组的子数组得到一个数组
func canChoose(groups [][]int, nums []int) bool {
	count := 0
	for i := 0; i < len(groups); i++ {
		arr := groups[i]
		flag := false
		for count+len(arr) <= len(nums) {
			if judge(arr, nums[count:count+len(arr)]) == true {
				count = count + len(arr)
				flag = true
				break
			} else {
				count = count + 1
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

func judge(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
