package main

func main() {

}

// leetcode1744_你能在你最喜欢的那天吃到你最喜欢的糖果吗？
func canEat(candiesCount []int, queries [][]int) []bool {
	arr := make([]int, len(candiesCount)+1)
	for i := 1; i <= len(candiesCount); i++ {
		arr[i] = arr[i-1] + candiesCount[i-1]
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		a := queries[i][0] // 第几类
		b := queries[i][1] // 第几天，到第n天共n+1天
		c := queries[i][2] // 每天最多吃c颗
		total := arr[a+1]
		if total <= b { // 最少一天一颗
			continue
		}
		if c*(b+1) <= arr[a] { // 最多每天吃c颗
			continue
		}
		res[i] = true
	}
	return res
}
