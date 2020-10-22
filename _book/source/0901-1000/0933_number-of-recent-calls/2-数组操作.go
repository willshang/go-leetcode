package main

func main() {

}

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
	start := t - 3000
	for len(r.arr) > 0 && r.arr[0] < start {
		r.arr = r.arr[1:]
	}
	return len(r.arr)
}
