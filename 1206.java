class Skiplist {
    static final int MAX_LEVEL = 32;
    static final double P_FACTOR = 0.25;
    private SkiplistNode head;
    private int level;
    private Random random; 

    public Skiplist() {
        this.head = new SkiplistNode(-1, MAX_LEVEL);
        this.level = 0;
        this.random = new Random();
    }
    
    public boolean search(int target) {
        SkiplistNode cur = this.head;
        for(int i=level-1; i >= 0; i--){
            while(cur.forward[i] != null && cur.forward[i].val < target){
                cur = cur.forward[i];
            }
        }
        cur = cur.forward[0];
        if(cur != null && cur.val == target){
            return true;
        }
        return false;
    }
    
    public void add(int num) {
        SkiplistNode[] update = new SkiplistNode[MAX_LEVEL];
        Arrays.fill(update, head);
        SkiplistNode cur = this.head;
        for(int i = level-1; i>=0; i--){
            while(cur.forward[i] != null && cur.forward[i].val < num){
                cur = cur.forward[i];
            }
            update[i] = cur;
        }
        int lv = randomLevel();
        level = Math.max(level, lv);
        SkiplistNode newNode = new SkiplistNode(num, lv);
        for(int i=0; i < lv; i++){
            newNode.forward[i] = update[i].forward[i];
            update[i].forward[i] = newNode;
        }
    }
    
    public boolean erase(int num) {
        SkiplistNode[] update = new SkiplistNode[MAX_LEVEL];
        SkiplistNode cur = this.head;
        for(int i=level-1; i >=0; i--){
            while(cur.forward[i] != null && cur.forward[i].val < num){
                cur = cur.forward[i];
            }
            update[i] = cur;
        }
        cur = cur.forward[0];
        if(cur == null || cur.val != num){
            return false;
        }
        for(int i=0; i<level; i++){
            if(update[i].forward[i] != cur){
                break;
            }
            update[i].forward[i] = cur.forward[i];
        }
        while(level > 1 && head.forward[level -1 ] == null){
            level--;
        }
        return true;
    }
    private int randomLevel(){
        int lv = 1;
        while(random.nextDouble() < P_FACTOR && lv < MAX_LEVEL){
            lv++;
        }
        return lv;
    }

    class SkiplistNode{
        int val;
        SkiplistNode[] forward;

        public SkiplistNode(int val, int maxLevel){
            this.val = val;
            this.forward = new SkiplistNode[maxLevel];
        }
    }
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * Skiplist obj = new Skiplist();
 * boolean param_1 = obj.search(target);
 * obj.add(num);
 * boolean param_3 = obj.erase(num);
 */