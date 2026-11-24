package main

import (
	"fmt"
	"sort"
)

// SubsetSums returns the sums of all 2^n subsets, sorted ascending.
// Each element is either taken or skipped — the recursion tree's leaves
// are exactly the subsets.
func SubsetSums(nums []int) []int {
	out := make([]int, 0, 1<<len(nums))
	var walk func(i, sum int)
	walk = func(i, sum int) {
		if i == len(nums) {
			out = append(out, sum)
			return
		}
		walk(i+1, sum+nums[i]) // take nums[i]
		walk(i+1, sum)         // skip nums[i]
	}
	walk(0, 0)
	sort.Ints(out)
	return out
}

func main() {
	fmt.Println(SubsetSums([]int{3, 1, 2}))
}
