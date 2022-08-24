/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	res := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			res = append(res, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, value := range res {
		curNode.Right = &TreeNode{Val: value}
		curNode = curNode.Right
	}
	return dummyNode.Right
}