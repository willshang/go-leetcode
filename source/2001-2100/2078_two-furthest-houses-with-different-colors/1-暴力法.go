package main

func main() {

}

// leetcode2078_两栋颜色不同且距离最远的房子
func maxDistance(colors []int) int {
	res := 0
	for i := 0; i < len(colors); i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] != colors[j] {
				res = max(res, j-i)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
