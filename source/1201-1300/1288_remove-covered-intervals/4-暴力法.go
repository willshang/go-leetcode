package main

func main() {

}

func removeCoveredIntervals(intervals [][]int) int {
	res := len(intervals)
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals); j++ {
			if i != j && intervals[j][0] <= intervals[i][0] &&
				intervals[i][1] <= intervals[j][1] {
				res--
				break
			}
		}
	}
	return res
}
