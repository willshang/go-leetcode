package main

func main() {

}

// 剑指OfferII093.最长斐波那契数列
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[arr[i]] = true
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count := 2
			a, b := arr[i], arr[j]
			for m[a+b] == true {
				count++
				a, b = b, a+b
			}
			if count > res && count > 2 {
				res = count
			}
		}
	}
	return res
}
