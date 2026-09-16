func getConcatenation(nums []int) []int {
	numsLength := len(nums)
    ans := make([]int, numsLength * 2)
	for idx, val := range(nums) {
		ans[idx] = val
		ans[idx + numsLength] = val
	}
	return ans
}
