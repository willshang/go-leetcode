package main

import "fmt"

func main() {
	fmt.Println('a')
	fmt.Println('1')
}

func squareIsWhite(coordinates string) bool {
	// a => 97  1 => 49
	return (coordinates[0]+coordinates[1])%2 != 0
}
