package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scan.Scan()
		str := scan.Text()
		length := len(str)
		if length%8 != 0 {
			for i := 0; i < 8-length%8; i++ {
				str = str + string('0')
			}
		}
		for i := 0; i < len(str)/8; i++ {
			fmt.Println(str[i*8 : i*8+8])
		}
	}
}
