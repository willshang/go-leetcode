package main

func main() {

}

// leetcode2115_从给定原材料中找到所有可以做出的菜
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	res := make([]string, 0)
	arr := make(map[string][]string)
	inDegree := make(map[string]int)
	for i := 0; i < len(recipes); i++ {
		a := recipes[i]
		for j := 0; j < len(ingredients[i]); j++ {
			b := ingredients[i][j] // b=>a 原材料=>菜
			arr[b] = append(arr[b], a)
			inDegree[a]++ // 菜的入度+1
		}
	}
	for len(supplies) > 0 {
		b := supplies[0]
		supplies = supplies[1:]
		for i := 0; i < len(arr[b]); i++ {
			a := arr[b][i]
			inDegree[a]--
			if inDegree[a] == 0 { // 入度=0，代表该菜的原材料都有
				res = append(res, a)
				supplies = append(supplies, a)
			}
		}
	}
	return res
}
