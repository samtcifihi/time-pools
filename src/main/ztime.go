package main

// dwells, breathers, trices, lulls
type ztime [4]byte

func (z *ztime) dec() bool {
	timeRemaining := true

	if z[3] != 0 {
		z[3]--
	} else if z[2] != 0 {
		z[2]--
		z[3] = 11
	} else if z[1] != 0 {
		z[1]--
		z[2], z[3] = 11, 11
	} else if z[0] != 0 {
		z[0]--
		z[1], z[2], z[3] = 11, 11, 11
	} else {
		timeRemaining = false
	}

	return timeRemaining
}

// func inc
func (z *ztime) inc() {
	if z[3] != 11 {
		z[3]++
	} else if z[2] != 11 {
		z[2]++
		z[3] = 0
	} else if z[1] != 11 {
		z[1]++
		z[2], z[3] = 0, 0
	} else {
		z[0]++
		z[1], z[2], z[3] = 0, 0, 0
	}
}
