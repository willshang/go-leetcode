package main

import "sort"

func main() {

}

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1] // 相同情况下，防御力从小到大
		}
		return properties[i][0] > properties[j][0] // 攻击力从大到小
	})
	res := 0
	stack := make([]int, 0)
	for i := 0; i < len(properties); i++ {
		// 出栈：小于当前防御力的出栈
		for len(stack) > 0 && properties[stack[len(stack)-1]][1] <= properties[i][1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 { // 栈顶存在大于当前数，弱角色+1
			res++
		}
		stack = append(stack, i)
	}
	return res
}
