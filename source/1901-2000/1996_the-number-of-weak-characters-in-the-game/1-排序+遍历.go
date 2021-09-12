package main

import "sort"

func main() {

}

// leetcode1996_游戏中弱角色的数量
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1] // 相同情况下，防御力从小到大
		}
		return properties[i][0] > properties[j][0] // 攻击力从大到小
	})
	res := 0
	maxValue := 0                          // 最大的防御力
	for i := 0; i < len(properties); i++ { // 攻击力从大到小
		// 注意：攻击力相同的情况，防御力是从小到大的，此时会更新maxValue
		if properties[i][1] < maxValue { // 当前面出现防御力大于当前防御力=>+1
			res++
		} else {
			maxValue = properties[i][1] // 更新防御力
		}
	}
	return res
}
