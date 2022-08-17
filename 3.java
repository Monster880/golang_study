class Solution {
    public int lengthOfLongestSubstring(String s) {
        Set<Character> set = new HashSet<Character>();
        int n = s.length();
        int right = -1, res = 0;
        for(int i=0; i < n; i++){
            if(i != 0){
                set.remove(s.charAt(i-1));
            }
            while(right + 1 < n && !set.contains(s.charAt(right + 1))){
                set.add(s.charAt(right + 1));
                right++;
            }
            res = Math.max(res, right - i + 1);
        }
        return res;
    }
}
