package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: the root of binary tree
 * @return: the length of the longest consecutive sequence path
 */

var (
	final = 1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	traverse(root, 0)
	return final
}

func traverse(n *TreeNode, l int) {
	if n.Left != nil {
		if n.Left.Val == n.Val+1 {
			if final < l+1 {
				final = l + 1
			}
			traverse(n.Left, l+1)
		} else {
			traverse(n.Left, 0)
		}
	}
	if n.Right != nil {
		if n.Right.Val == n.Val+1 {
			if final < l+1 {
				final = l + 1
			}
			traverse(n.Right, l+1)
		} else {
			traverse(n.Right, 0)
		}
	}
}

//Input {1,#,3,2,4,#,#,#,5}
/**
		1
#				3
			  2   4
*/
