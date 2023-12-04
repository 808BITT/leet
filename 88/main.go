package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// insert
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	// sort
	sort.Ints(nums1)

	// print
	for i := 0; i < m+n; i++ {
		fmt.Print(nums1[i])
	}
	fmt.Println()
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3) // expect [1,2,2,3,5,6]
	merge([]int{1}, 1, []int{}, 0)                       // expect [1]
	merge([]int{0}, 0, []int{1}, 1)                      // expect [1]
}
