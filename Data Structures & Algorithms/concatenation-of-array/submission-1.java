class Solution {
    public int[] getConcatenation(int[] nums) {
        List<Integer> newNums = new ArrayList();
        int numsIndex = 0;
        while(newNums.size() < nums.length * 2){
            newNums.add(nums[numsIndex]);
            if(numsIndex == nums.length - 1){
                numsIndex = 0;
            } else {
                numsIndex++;
            }
        }
        return newNums.stream().mapToInt(Integer::intValue).toArray();
    }
}