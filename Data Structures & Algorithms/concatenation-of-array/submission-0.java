class Solution {
    public int[] getConcatenation(int[] nums) {
        int[] newNums = new int[nums.length * 2];
        for(int index = 0; index < nums.length; index++){
            newNums[index] = nums[index];
            newNums[index + nums.length] = nums[index];
        }
        return newNums;
    }
}