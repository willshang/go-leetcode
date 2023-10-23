package main

func main() {

}

// leetcode1663_具有给定数值的最小字符串
func getSmallestString(n int, k int) string {
	arr := make([]byte, n)
	k = k - n
	for i := n - 1; i >= 0; i-- {
		if k > 25 {
			arr[i] = 'z'
			k = k - 25
		} else {
			arr[i] = byte('a' + k)
			k = 0
		}
	}
	return string(arr)
}
