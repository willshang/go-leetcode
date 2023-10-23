package main

func main() {

}

func reinitializePermutation(n int) int {
	res := 0
	target := make([]int, n)
	perm := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	copy(target, perm)
	for {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		res++
		flag := true
		for i := 0; i < n; i++ {
			if arr[i] != target[i] {
				flag = false
				break
			}
		}
		if flag == true {
			break
		}
		copy(perm, arr)
	}
	return res
}
