package main

import (
	"fmt"
	"math"
)

//titleToNumber: rune is 64 onwards for letter A...
func titleToNumber(s string) int {
	length := len(s)
	sum := 0
	for i := 0; i < length; i++ {
		runeEquivalent := int(s[i]) - 64
		powerPlacing := math.Pow(26, float64(length-i-1))
		sum += int(powerPlacing) * runeEquivalent
	}
	return sum
}

func encodeChar(i int32) string {
	return string(i + 64)
}

func numberToTitle(i int32) string {
	result := ""
	var r, q int32
	ok := true

	for ok == true {
		q = i / 26
		r = i % 26
		result = encodeChar(r) + result
		if q <= 26 {
			ok = false
			if q != 0 {
				result = encodeChar(q) + result
			}
		}
		i = q
	}

	return result
}

func main() {

	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ABC"))
	fmt.Println(numberToTitle(2))
	fmt.Println(numberToTitle(28))
	fmt.Println(numberToTitle(73111))
}
