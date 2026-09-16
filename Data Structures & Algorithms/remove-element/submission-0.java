class Solution {
    public int removeElement(int[] nums, int val) {
        int differentValues = 0;
        for(int index = 0; index < nums.length; index++){
            if(nums[index] != val){
                nums[differentValues] = nums[index];
                differentValues++;
            }
        }
        return differentValues;
    }
}