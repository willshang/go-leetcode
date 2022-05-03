package main

func main() {

}

// leetcode2240_买钢笔和铅笔的方案数
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	res := int64(0)
	for i := 0; i <= total/cost1; i++ {
		res = res + int64((total-cost1*i)/cost2) + 1 // 可以买0支铅笔，次数+1
	}
	return res
}
