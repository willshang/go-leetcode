package main

func main() {

}

func addNegabinary(arr1 []int, arr2 []int) []int {
	res := make([]int, 1005)
	last := 1004
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry = carry + arr1[i]
			i--
		}
		if j >= 0 {
			carry = carry + arr2[j]
			j--
		}
	}
	return res
}
