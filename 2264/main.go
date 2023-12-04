package main

import (
	"strconv"
)

func largestGoodInteger(num string) string {
	highest := -1
	// find all instances where 3 digits repeat
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			// convert to int
			n, _ := strconv.Atoi(num[i : i+3])
			if n > highest {
				highest = n
			}
		}
	}
	if highest == -1 {
		return ""
	}

	if highest == 0 {
		return "000"
	}

	return strconv.Itoa(highest)
}

func main() {
	ans := largestGoodInteger("334111434555234")
	println(ans)
}
