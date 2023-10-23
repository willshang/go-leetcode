package main

import "strconv"

func main() {

}

func splitString(s string) bool {
	for i := 0; i < len(s); i++ {
		value, _ := strconv.Atoi(s[:i+1])
		if s[i+1:] != "" && dfs(s[i+1:], value-1) == true {
			return true
		}
	}
	return false
}

func dfs(s string, target int) bool {
	value, _ := strconv.Atoi(s)
	if s == "" || value == target {
		return true
	}
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[:i+1])
		if v < target {
			continue
		}
		if v > target {
			return false
		}
		if v == target {
			if dfs(s[i+1:], target-1) == true {
				return true
			}
			return false
		}
	}
	return false
}
