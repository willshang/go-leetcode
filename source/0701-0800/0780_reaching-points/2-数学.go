package main

func main() {

}

// leetcode780_到达终点
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx = tx % ty
		} else {
			ty = ty % tx
		}
	}
	if tx == sx { // (x,y) => (x, kx+y)
		return ty >= sy && (ty-sy)%sx == 0
	}
	if ty == sy { // (x,y) => (x+ky,y)
		return tx >= sx && (tx-sx)%sy == 0
	}
	return false
}
