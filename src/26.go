package main

import "fmt"

func removeDuplicates(nums []int) int {
	insIndex := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		if i == 0 {
			nums[insIndex] = nums[i]
			insIndex++
			continue
		}
		if nums[i] != nums[i-1] {
			nums[insIndex] = nums[i]
			insIndex++
		}
	}

	nums = nums[:insIndex]
	return len(nums)
}

func main() {
	val1 := removeDuplicates([]int{1, 1, 2})                      // expect 2
	val2 := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) // expect 5
	fmt.Println(val1)
	fmt.Println(val2)
}
