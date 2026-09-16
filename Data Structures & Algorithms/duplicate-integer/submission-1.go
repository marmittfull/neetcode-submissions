func hasDuplicate(nums []int) bool {
    seemElements := make(map[int]int)

	for _, val := range nums {
		_, exists := seemElements[val]
		if exists {
			return true
		}
		seemElements[val] = val
	}
	return false
}
