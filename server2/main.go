package main

import (
	"fmt"

	KEY "./keyGenerator"
)

func main() {
	str := KEY.KeyGen([]uint{2, 2, 1, 1, 1, 2})
	fmt.Println(str)
}
