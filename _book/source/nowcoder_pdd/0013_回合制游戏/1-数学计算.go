package main

import "fmt"

func getCount(hp, normalAttack, buffedAttack int) int {
	res := 0
	if normalAttack*2 < buffedAttack {
		res = res + 2*(hp/buffedAttack)
		hp = hp - hp/buffedAttack*buffedAttack
		if hp > 0 {
			if hp <= normalAttack {
				res = res + 1
			} else {
				res = res + 2
			}
		}
	} else {
		res = res + hp/normalAttack
		if hp%normalAttack > 0 {
			res = res + 1
		}
	}
	return res
}

func main() {
	var hp, normalAttack, buffedAttack int
	for {
		// 三行数字用fmt.Scan
		a, _ := fmt.Scan(&hp, &normalAttack, &buffedAttack)
		if a == 0 {
			break
		}
		fmt.Println(getCount(hp, normalAttack, buffedAttack))
	}
}
