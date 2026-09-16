class Solution {
    public boolean isValid(String s) {
        char[] chairs = s.toCharArray();
        Stack<Character> bracketStack = new Stack<>();
        for(int index = 0; index < chairs.length; index++){
            if(chairs[index] == '(' || chairs[index] == '[' || chairs[index] == '{'){
                bracketStack.push(chairs[index]);
            } else if(!bracketStack.isEmpty() && Arrays.asList("()", "[]", "{}").contains("" + bracketStack.peek() + chairs[index])){
                bracketStack.pop();
            } else {
                return false;
            }
        }
        return bracketStack.isEmpty();
    }
}