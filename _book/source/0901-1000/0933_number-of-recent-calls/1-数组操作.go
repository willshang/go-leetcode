package main

func main() {

}

// leetcode933_最近的请求次数
type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		arr: make([]int, 0),
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.arr = append(r.arr, t)
	res := 1
	for i := len(r.arr) - 2; i >= 0; i-- {
		if t-r.arr[i] <= 3000 {
			res++
		} else {
			r.arr = r.arr[i+1:]
			break
		}
	}
	return res
}
