// This package util covers common utilities
package util

import (
	"math"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Map of alphabet characters by given integer.
var alphabetChar = map[int]string{
	0:  "",
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

// Map of integers by given alphabet character.
var alphabetInt = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

// Convert float to integer and make a round down
func toIntRoundDown(f float64) int {
	if f < 0 {
		return int(f - 1.0)
	}
	return int(f)
}

// Get calculated index of 2nd character from given pointer integer
func secondCharItos(x int) int {
	var root, threshold float64
	for i := 0; i < 4; i++ {
		log.Debug("i:" + strconv.Itoa(i))
		f := float64(i)
		log.Debug("pangkat: " + strconv.Itoa(int(math.Pow(26, f))))
		t := math.Pow(26, f) + float64((i-1)*26)
		log.Debug("Z: " + strconv.Itoa(int(t)))
		if float64(x) >= t {
			root = f
			threshold = t
		}
	}
	log.Debug("Root: " + strconv.Itoa(int(root)))
	log.Debug("Power: " + strconv.Itoa(int(math.Pow(26, root))))
	//y := float64(x) - math.Pow(26, root)
	y := float64(x) - math.Abs(threshold) + 26
	d := float64(y / 26)
	return toIntRoundDown(d)
}

// Get calculated index of 3rd character from given pointer integer
func thirdCharItos(x int) int {
	return (x % 26) + 1
}

// Get representing characters by given pointer integer
// I.e. Giving pointer integer 27 (the 28th column) will give you XLSX cell
// characters: AB
func CharsOfPointerInt(i int) string {
	var c1, c2 string
	// Currently support only up to 2 characters
	if i1 := secondCharItos(i); i1 > -1 {
		c1 = alphabetChar[i1]
	}
	if i2 := thirdCharItos(i); i2 > -1 {
		c2 = alphabetChar[i2]
	}
	return c1 + c2
}

// Provide representing integer from given 1st pointer character
func firstCharStoi(x string) int {
	i := -1
	if a, exists := alphabetInt[x]; exists {
		log.Debug("x: " + x + ", a: " + strconv.Itoa(a))
		y := int(math.Pow(26, 2))
		i = y + (675 * (a - 1)) + (a - 1)
		log.Debug("y:" + strconv.Itoa(y))
	}
	log.Debug("first: " + strconv.Itoa(i))
	return i
}

// Provide representing integer from given 2nd pointer character
func secondCharStoi(x string) int {
	i := -1
	if a, exists := alphabetInt[x]; exists {
		log.Debug("x: " + x + ", a: " + strconv.Itoa(a))
		i = 26 * a
	}
	log.Debug("second: " + strconv.Itoa(i))
	return i
}

// Provide representing integer from given 3rd pointer character
func thirdCharStoi(x string) int {
	i := -1
	if a, exists := alphabetInt[x]; exists {
		log.Debug("x: " + x + ", a: " + strconv.Itoa(a))
		i = a - 1
	}
	log.Debug("third: " + strconv.Itoa(i))
	return i
}

// Get representing integer by given pointer characters
// I.e. Giving pointer chars of XLSX cell "AB" (the AB column) will give you: 27
func IntOfPointerChars(s string) int {
	x := -1
	a := strings.Split(strings.ToUpper(strings.Trim(s, " ")), "")
	log.Debug("len: " + strconv.Itoa(len(a)))
	// Support up to 3 characters
	if len(a) == 1 {
		log.Debug("1 char")
		x = thirdCharStoi(a[0])
	} else if len(a) == 2 {
		log.Debug("2 chars")
		i := secondCharStoi(a[0])
		j := thirdCharStoi(a[1])
		if i >= -1 && j >= -1 {
			x = i + j
		}
	} else if len(a) == 3 {
		log.Debug("3 chars")
		i := firstCharStoi(a[0])
		j := secondCharStoi(a[1])
		k := thirdCharStoi(a[2])
		if i >= -1 && j >= -1 && k >= -1 {
			x = i + j + k
		}
	} else {
		log.Error("Character's length is not support")
	}
	log.Debug("result: " + strconv.Itoa(x))
	return x
}
