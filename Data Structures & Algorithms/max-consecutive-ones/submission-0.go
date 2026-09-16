func findMaxConsecutiveOnes(nums []int) int {
	var maxConsecutivesOnes int = 0
	var currentConsecutivesOnes int = 0
	for idx, val := range(nums) {
		if(val == 1){
			currentConsecutivesOnes++
		} 
		if(val != 1 || idx == len(nums) - 1) {
			if(currentConsecutivesOnes > maxConsecutivesOnes){
				maxConsecutivesOnes = currentConsecutivesOnes
			}
			currentConsecutivesOnes = 0
		}
	}
	return maxConsecutivesOnes
}
