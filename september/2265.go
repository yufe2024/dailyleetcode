package september

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归检查每个节点
//
//	func averageOfSubtree(root *TreeNode) int {
//		ans := 0
//		if root == nil {
//			return ans
//		}
//		//这里需要一个值来存储所有子树的加和，每个节点都需要
//		//先遍历一遍放到数组里再用数组做怎么样
//		//不行，无法根据遍历的结果唯一确定一棵树
//		//那就左子树右子树分别处理
//		//定义一个全局的解吧
//
// }
// 题解写的很好
func averageOfSubtree(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftSize := dfs(node.Left)
		rightSum, rightSize := dfs(node.Right)
		Size := leftSize + rightSize + 1
		Sum := leftSum + rightSum + node.Val
		if Size > 0 && Sum/Size == node.Val {
			ans++
		}
		return Sum, Size
	}
	dfs(root)
	return ans
}
