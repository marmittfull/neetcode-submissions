class Solution {
    public int removeDuplicates(int[] nums) {
        int uniqueNumbers = 0;
        for(int index = 0; index < nums.length; index++){
            if(index == 0 || nums[index] != nums[index - 1]){
                nums[uniqueNumbers] = nums[index];
                uniqueNumbers++;
            }
        }
        return uniqueNumbers;
    }
}