package main

import (
	"fmt"
	"github.com/imroc/biu"
)

func main() {
	// fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	fmt.Println(biu.Uint32ToBinaryString(num))
	num = ((num & 0xffff0000) >> 16) | ((num & 0x0000ffff) << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}
