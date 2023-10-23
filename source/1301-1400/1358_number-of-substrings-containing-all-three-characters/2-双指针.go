package main

func main() {

}

func numberOfSubstrings(s string) int {
	res := 0
	n := len(s)
	arr := [3]int{}
	left := 0
	right := -1
	var value int
	for left = 0; left < n; left++ {
		for right < n && (arr[0] == 0 || arr[1] == 0 || arr[2] == 0) {
			right++
			if right == n {
				break
			}
			value = int(s[right] - 'a')
			arr[value]++
		}
		res = res + n - right
		value = int(s[left] - 'a')
		arr[value]--
	}
	return res
}
