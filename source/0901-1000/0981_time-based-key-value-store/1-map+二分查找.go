package main

func main() {

}

// leetcode981_基于时间的键值存储
type Node struct {
	timestamp int
	str       string
}

type TimeMap struct {
	m map[string][]Node
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]Node)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Node{
		timestamp: timestamp,
		str:       value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.m[key]
	n := len(arr)
	if n == 0 || (timestamp < arr[0].timestamp) {
		return ""
	}
	if timestamp >= arr[n-1].timestamp {
		return arr[n-1].str
	}
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if arr[mid].timestamp == timestamp {
			return arr[mid].str
		} else if arr[mid].timestamp < timestamp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return arr[left].str
}
