func containsDuplicate(nums []int) bool {

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {

			if nums[i] == nums[j] && i != j {
				return true
			}
		}
	}

	return false
}