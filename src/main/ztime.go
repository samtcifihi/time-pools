package main

import (
	"strconv"
)

// dwells, breathers, trices, lulls
type ztime [4]int

// Format formats a ztime according to a provided format string.
// Accepted tokens are %d, %b, %t, and %l.
// %% escapes the second %
func (z *ztime) Format(formatString string) string {
	output := ""

	// separate formatString into tokens
	tokens := []string{}
	skipNext := false
	for i, v := range formatString {
		if skipNext == false {
			if (v != '%') || (len(formatString) < i+2) {
				tokens = append(tokens, string(v))
			} else {
				tokens = append(tokens, formatString[i:i+2])
				skipNext = true
			}
		} else {
			skipNext = false
		}
	}

	// evaluate tokens and assign to output
	for _, v := range tokens {
		if len(v) == 1 {
			output += v
		} else {
			switch v {
			case "%d":
				output += dtoz(z[0])
			case "%b":
				output += dtoz(z[1])
			case "%t":
				output += dtoz(z[2])
			case "%l":
				output += dtoz(z[3])
			case "%%":
				output += "%"
			}
		}
	}

	return output
}

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

// Sum adds two ztimes, returning the result and any carry
// carry will be 0 unless there is an overflow
func Sum(a ztime, b ztime) (ztime, int) {
	var z ztime
	var c int

	for i := 3; i >= 0; i-- {
		// a + b + c mod 0z10
		// c = a + b + c - (a + b + c mod 0z10)
		z[i] = (a[i] + b[i] + c) % 12
		c = (a[i] + b[i] + c) / 12
	}

	return z, c
}

// dtoz converts a decimal number [0, 11] to a dozenal digit
// it has poor to no error reporting
func dtoz(i int) string {
	if i < 10 && i >= 0 {
		return strconv.Itoa(i)
	}

	switch i {
	case 10:
		return "\u218a"
	case 11:
		return "\u218b"
	default:
		return ""
	}
}
