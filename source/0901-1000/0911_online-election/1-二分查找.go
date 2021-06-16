package main

func main() {

}

type TopVotedCandidate struct {
	result []int // 存放对应时间的候选人
	times  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	arr := make([]int, n+1)
	maxCount := 0
	index := persons[0]
	res := make([]int, n)
	for i := 0; i < n; i++ {
		id := persons[i]
		arr[id]++
		if arr[id] >= maxCount {
			maxCount = arr[id]
			index = id
		}
		res[i] = index
	}
	return TopVotedCandidate{
		result: res,
		times:  times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	left, right := 0, len(this.times)
	for left < right {
		mid := left + (right-left)/2
		if this.times[mid] == t {
			return this.result[mid]
		} else if this.times[mid] > t {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return this.result[left-1]
}
