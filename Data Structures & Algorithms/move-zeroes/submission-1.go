func moveZeroes(nums []int) {
	var lastValidPosition int
	for idx, val := range nums {
		if val > 0 {
			nums[lastValidPosition] = val

			if lastValidPosition != idx {
				nums[idx] = 0
			}

			lastValidPosition++
		}
	}
}
