package main

func main() {

}

// leetcode898_子数组按位或操作
func subarrayBitwiseORs(arr []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
		temp := 0
		for j := i - 1; j >= 0; j-- {
			if temp|arr[i] == temp { // 都为1，进行下去无意义，避免超时
				break
			}
			temp = temp | arr[j]
			m[temp|arr[i]] = true
		}
	}
	return len(m)
}
