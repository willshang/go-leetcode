package main

func main() {

}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx == ty {
			break
		}
		if tx > ty {
			// (tx,ty) => (tx-ty,ty)
			if ty > sy {
				tx = tx % ty
			} else {
				return (tx-sx)%sy == 0
			}
		} else {
			if tx > sx {
				ty = ty % tx
			} else {
				return (ty-sy)%sx == 0
			}
		}
	}
	return tx == sx && ty == sy
}
