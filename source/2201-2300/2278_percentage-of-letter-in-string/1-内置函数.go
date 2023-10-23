package main

import "strings"

func main() {

}

func percentageLetter(s string, letter byte) int {
	return 100 * strings.Count(s, string(letter)) / len(s)
}
