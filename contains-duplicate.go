package main

func containsDuplicate(nums []int) bool {

	duplicates := make(map[int]struct{})

	for _, i := range nums {

		if _, ok := duplicates[i]; ok {
			return true
		} else {
			duplicates[i] = struct{}{}
		}
	}

	return false
}

// Runtime: 78 ms, faster than 89.12% of Go online submissions for Contains Duplicate.
// Memory Usage: 8.6 MB, less than 77.05% of Go online submissions for Contains Duplicate.

// func containsDuplicate(nums []int) bool {

// 	for i := 0; i < len(nums); i++ {

// 		for j := 0; j < len(nums); j++ {

// 			if nums[i] == nums[j] && i != j {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// Runtime: 2233 ms, faster than 5.02% of Go online submissions for Contains Duplicate.
// Memory Usage: 8.2 MB, less than 94.94% of Go online submissions for Contains Duplicate.
