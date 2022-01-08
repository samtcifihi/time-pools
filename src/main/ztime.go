package main

// dwells, breathers, trices, lulls
type ztime [4]byte

func (z *ztime) Dec() bool {
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
func (z *ztime) Inc() {
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

// first digit is dwells regardless of input length
func (z *ztime) Set(time string) {
	for x := 3; x >= 0; x-- {
		if x < len(time) {
			switch time[x] {
			case '0':
				z[x] = 0
			case '1':
				z[x] = 1
			case '2':
				z[x] = 2
			case '3':
				z[x] = 3
			case '4':
				z[x] = 4
			case '5':
				z[x] = 5
			case '6':
				z[x] = 6
			case '7':
				z[x] = 7
			case '8':
				z[x] = 8
			case '9':
				z[x] = 9
			case 'a':
				z[x] = 10
			case 'b':
				z[x] = 11
			}
		}
	}
}
