package main

func main() {

}

func rand10() int {
	for {
		a := rand7()
		b := rand7()
		target := a + (b-1)*7
		if target <= 40 { // 1-49
			return target%10 + 1
		}
		a = rand7()
		b = target - 40
		target = a + (b-1)*7 // 1-63
		if target <= 60 {
			return target%10 + 1
		}
		a = rand7()
		b = target - 60
		target = a + (b-1)*7 // 1-21
		if target <= 20 {
			return target%10 + 1
		}
	}
}
