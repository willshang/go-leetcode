package main

func main() {

}

// leetcode649_Dota2参议院
func predictPartyVictory(senate string) string {
	r, d := make([]int, 0), make([]int, 0)
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
