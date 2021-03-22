package main

func main() {

}

// leetcode1773_统计匹配检索规则的物品数量
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for i := 0; i < len(items); i++ {
		if ruleKey == "type" && items[i][0] == ruleValue {
			res++
		} else if ruleKey == "color" && items[i][1] == ruleValue {
			res++
		} else if ruleKey == "name" && items[i][2] == ruleValue {
			res++
		}
	}
	return res
}
