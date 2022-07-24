class Solution {
    public boolean sequenceReconstruction(int[] nums, int[][] sequences) {
        int n = nums.length;
        int[] indegrees = new int[n+1];
        Set<Integer>[] graph = new Set[n+1];
        for(int i=1; i<=n; i++){
            graph[i] = new HashSet<Integer>();
        }
        for(int[] sequence : sequences){
            int size = sequence.length;
            for(int i=1; i<size; i++){
                int prev = sequence[i-1], next = sequence[i];
                if(graph[prev].add(next)){
                    indegrees[next]++;
                }
            }
        }
        Queue<Integer> q = new LinkedList<>();
        for(int i=1; i<=n; i++){
            if(indegrees[i] == 0){
                q.offer(i);
            }
        }
        while(!q.isEmpty()){
            if(q.size() > 1){
                return false;
            }
            int num = q.poll();
            Set<Integer> set = graph[num];
            for(int next : set){
                indegrees[next]--;
                if(indegrees[next] == 0){
                    q.offer(next);
                }
            }
        }
        return true;
    }
}