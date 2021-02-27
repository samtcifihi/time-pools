package main

import (
	"fmt"
)

var znum = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "\u218A", "\u218B"}

func main() {
	fmt.Println("znum test:")

	for x := 0; x <= 11; x++ {
		fmt.Println(znum[x])
	}

	testTime := ztime{0, 0, 1, 3}

	fmt.Println("Initial Time:", testTime)

	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)
	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)
	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)
	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)
	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)
	fmt.Println("Time left?", testTime.dec(),
		"\nCurrent Time Remaining", testTime)

	fmt.Println("Increment Test")
	testTime.inc()
	fmt.Println(testTime)
	testTime.inc()
	fmt.Println(testTime)
	testTime.inc()
	fmt.Println(testTime)
	testTime.inc()
	fmt.Println(testTime)
	testTime.inc()
	fmt.Println(testTime)
}
