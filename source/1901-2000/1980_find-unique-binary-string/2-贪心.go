package main

func main() {

}

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	res := ""
	// 康托对角线
	// 只要和nums[i][i]不同，构造出的串就和所有的串都不同
	for i := 0; i < n; i++ {
		if nums[i][i] == '0' {
			res = res + "1"
		} else {
			res = res + "0"
		}
	}
	return res
}
