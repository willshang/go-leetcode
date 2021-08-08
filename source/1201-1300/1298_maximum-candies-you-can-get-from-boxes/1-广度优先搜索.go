package main

func main() {

}

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	res := 0
	m := make(map[int]bool) // 访问过的盒子
	//keysM := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, initialBoxes...)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if m[node] == false {
			if status[node] == 1 { // 盒子是打开的状态
				m[node] = true
				res = res + candies[node]
				for i := 0; i < len(keys[node]); i++ { // 处理钥匙
					key := keys[node][i]
					status[key] = 1 // 打开盒子
					//if keysM[key] ==
				}
			}
		}
	}
	return res
}
