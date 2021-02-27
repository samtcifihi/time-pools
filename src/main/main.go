package main

import (
	"fmt"
)

var znum = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '\u218A', '\u218B'}

func main() {
	testTime := ztime{0, 1, 2, 3}

	fmt.Println("Remaining Time:", testTime)
}
