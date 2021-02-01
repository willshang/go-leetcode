package main

func main() {

}

// leetcode904_水果成篮
func totalFruit(tree []int) int {
	res := 0
	m := make(map[int]int)
	left := 0
	total := 0
	for i := 0; i < len(tree); i++ {
		target := tree[i]
		m[target]++
		for len(m) > 2 {
			m[tree[left]]--
			total--
			if m[tree[left]] == 0 {
				delete(m, tree[left])
			}
			left++
		}
		total++
		if total > res {
			res = total
		}
	}
	return res
}
