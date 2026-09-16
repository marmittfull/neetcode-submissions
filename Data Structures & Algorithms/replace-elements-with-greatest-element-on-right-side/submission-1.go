func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	rightMax := -1

	for idx := n - 1; idx >= 0; idx-- {
		ans[idx] = rightMax
		if(arr[idx] > rightMax) {
			rightMax = arr[idx]
		}
	}
	return ans
}
