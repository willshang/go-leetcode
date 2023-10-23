package main

func main() {

}

// leetcode1396_设计地铁系统
type Node struct {
	startName string
	startTime int
}

type SumInfo struct {
	count int // 总次数
	total int // 总时常
}

type UndergroundSystem struct {
	record map[int]Node
	sum    map[string]SumInfo
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		record: make(map[int]Node),
		sum:    make(map[string]SumInfo),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.record[id] = Node{
		startName: stationName,
		startTime: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	key := this.record[id].startName + "=>" + stationName
	node := this.sum[key]
	node.count++
	node.total = node.total + t - this.record[id].startTime
	this.sum[key] = node
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "=>" + endStation
	return float64(this.sum[key].total) / float64(this.sum[key].count)
}
