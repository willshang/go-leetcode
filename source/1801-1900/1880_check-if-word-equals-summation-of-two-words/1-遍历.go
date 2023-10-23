package main

func main() {

}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	a := 0
	b := 0
	c := 0
	for i := 0; i < len(firstWord); i++ {
		a = a*10 + int(firstWord[i]-'a')
	}
	for i := 0; i < len(secondWord); i++ {
		b = b*10 + int(secondWord[i]-'a')
	}
	for i := 0; i < len(targetWord); i++ {
		c = c*10 + int(targetWord[i]-'a')
	}
	return a+b == c
}
