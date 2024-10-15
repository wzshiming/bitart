package main

import (
	"fmt"

	"github.com/wzshiming/bitart"
)

func main() {
	for i := 0; i < 256; i++ {
		fmt.Printf("0b_%08b %s\n", i, bitart.Bits8(i))
	}
}
