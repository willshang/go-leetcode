package main

import "fmt"

func main() {
	first := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	second := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(first, second))
}

// leetcode599_两个列表的最小索引总和
func findRestaurant(list1 []string, list2 []string) []string {
	min := 2000
	res := make([]string, 0, 1000)
	for key1, value1 := range list1 {
		for key2, value2 := range list2 {
			if value1 == value2 {
				if min == key1+key2 {
					res = append(res, value1)
				}
				if min > key1+key2 {
					min = key1 + key2
					res = []string{value1}
				}
			}
		}
	}
	return res
}
