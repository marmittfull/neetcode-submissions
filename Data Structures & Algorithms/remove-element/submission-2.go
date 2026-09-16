func removeElement(nums []int, val int) int {
    var validElements int = 0
    for currentPosition := 0; currentPosition < len(nums); currentPosition++ {
        if(nums[currentPosition] == val){
            continue
        }
        nums[validElements] = nums[currentPosition]
        validElements++
    }
    return validElements
}
