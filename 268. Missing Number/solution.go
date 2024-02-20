package main

import (
	"fmt"
	"slices"
)

func missingNumber(nums []int) int {
  slices.Sort(nums)
	len := len(nums)
	for i := 0; i < len; i++ {
		if v := nums[i]; v != i {
			return i
     }
	}
	return len
}

func main() {
	var nums1 = []int{3,0,1}
	var nums2 = []int{0,1}
	var nums3 = []int{9,6,4,2,3,5,7,0,1}

	var err = false

	if v := missingNumber(nums1); v != 2 {
		fmt.Printf("Case 1 incorrect, instead got %d\n", v)
		err = true
	}
	if v := missingNumber(nums2); v != 2 {
		fmt.Printf("Case 2 incorrect, instead got %d\n", v)
    err = true
  }
	if v := missingNumber(nums3); v != 8 {
		fmt.Printf("Case 3 incorrect, instead got %d\n", v)
    err = true
  }

	if !err {
		fmt.Println("Success!")
	}
}