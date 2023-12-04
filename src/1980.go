package main

import "strconv"

func findDifferentBinaryString(nums []string) string {
	length := len(nums[0])
	// create a map of all the nums
	m := make(map[string]bool)
	for _, num := range nums {
		m[num] = true
	}
	// loop through all possible binary strings
	for i := 0; i < 1<<length; i++ {
		// convert to binary
		binary := strconv.FormatInt(int64(i), 2)
		// pad with 0s
		for len(binary) < length {
			binary = "0" + binary
		}
		if !m[binary] {
			return binary
		}
	}
	return ""
}

func main() {
	ans := findDifferentBinaryString([]string{"011", "100"})
	println(ans)
}
