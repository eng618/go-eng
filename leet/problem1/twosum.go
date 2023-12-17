package problem1

// TwoSum is a function that finds the target sum of two ints, from a provided slices of ints.
// https://leetcode.com/problems/two-sum/
//
// Given an slice of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1].
func TwoSum(nums []int, target int) []int {
	// key = value to look up and compare, value = key to return
	numMap := make(map[int]int)

	for k, v := range nums {
		// First add value to map if it does not already exist
		if _, ok := numMap[v]; !ok {
			numMap[v] = k
			if v+v == target {
				// we just added this number so it can't be the same
				continue
			}
		}

		// If the first value continue as we have nothing to compare to.
		if k == 0 {
			continue
		}

		// Compare values to see if map have needed value
		if val, ok := numMap[target-v]; ok {
			return []int{val, k}
		}
	}

	return []int{}
}
