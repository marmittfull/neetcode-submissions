func findMaxConsecutiveOnes(nums []int) int {
	var maxConsecutivesOnes int = 0
	var currentConsecutivesOnes int = 0
	for _, val := range(nums) {
		if(val == 1){
			currentConsecutivesOnes++
		} else {
			currentConsecutivesOnes = 0
		}
		if(currentConsecutivesOnes > maxConsecutivesOnes){
			maxConsecutivesOnes = currentConsecutivesOnes
		}
	}
	return maxConsecutivesOnes
}
