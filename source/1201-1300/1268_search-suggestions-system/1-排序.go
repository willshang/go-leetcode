package main

import "sort"

func main() {

}

// leetcode1268_搜索推荐系统
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, 0)
	for i := 0; i < len(searchWord); i++ {
		target := searchWord[:i+1]
		arr := make([]string, 0)
		for j := 0; j < len(products); j++ {
			if len(products[j]) >= len(target) && products[j][:i+1] == target {
				arr = append(arr, products[j])
			}
			if len(arr) == 3 {
				break
			}
		}
		res = append(res, arr)
	}
	return res
}
