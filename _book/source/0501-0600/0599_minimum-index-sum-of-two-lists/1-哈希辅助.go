package main

import "fmt"

func main() {
	first := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	second := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(first, second))
}

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i := range list1 {
		m[list1[i]] = i
	}
	min := 2000
	res := make([]string, 0, 1000)
	for key2, value := range list2 {
		if key1, ok := m[value]; ok {
			if min == key1+key2 {
				res = append(res, value)
			}
			if min > key1+key2 {
				min = key1 + key2
				res = []string{value}
			}
		}
	}
	return res
}
