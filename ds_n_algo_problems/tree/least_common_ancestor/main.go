// https://www.interviewbit.com/problems/least-common-ancestor/
package main

//Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

/**
 * @input A : Root pointer of the tree
 * @input B : Integer
 * @input C : Integer
 *
 * @Output Integer
 */
func lca(A *treeNode, B int, C int) int {
	v, _ := search(A, B, C)
	return v
}

func search(A *treeNode, B, C int) (int, int) {
	found := 0
	f := 0
	v := -1
	if B == A.value {
		found++
	}

	if C == A.value {
		found++
	}

	if found == 2 {
		return A.value, found
	}

	if found != 2 && A.left != nil {
		v, f = search(A.left, B, C)
		found += f
	}

	if found != 2 && A.right != nil {
		v, f = search(A.right, B, C)
		found += f
	}

	if found == 2 {
		if v == -1 {
			return A.value, found
		}
		return v, found
	}

	return -1, found
}
