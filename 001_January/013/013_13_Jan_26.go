/*
2196. Create Binary Tree From Descriptions | Medium

You are given a 2D integer array descriptions where descriptions[i] = [parenti, childi, isLefti] indicates that parenti is the parent of childi in a binary tree of unique values. Furthermore,

If isLefti == 1, then childi is the left child of parenti.
If isLefti == 0, then childi is the right child of parenti.
Construct the binary tree described by descriptions and return its root.

The test cases will be generated such that the binary tree is valid.

Example 1:

Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
Output: [50,20,80,15,17,19]
Explanation: The root node is the node with value 50 since it has no parent.
The resulting binary tree is shown in the diagram.

Example 2:

Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
Output: [1,2,null,null,3,4]
Explanation: The root node is the node with value 1 since it has no parent.
The resulting binary tree is shown in the diagram.


Constraints:

1 <= descriptions.length <= 10^4
descriptions[i].length == 3
1 <= parenti, childi <= 10^5
0 <= isLefti <= 1
The binary tree described by descriptions is valid.
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// Create map to quick finding node
	nodes := make(map[int]*TreeNode)

	children := make(map[int]bool)

	// 1. Created node for kept relations
	for _, desc := range descriptions {
		parentVal, childVal, isLeft := desc[0], desc[1], desc[2]

		// If not have this node inside, Create a new one
		if nodes[parentVal] == nil {
			nodes[parentVal] = &TreeNode{Val: parentVal}
		}
		if nodes[childVal] == nil {
			nodes[childVal] = &TreeNode{Val: childVal}
		}

		// Connected the node
		if isLeft == 1 {
			nodes[parentVal].Left = nodes[childVal]
		} else {
			nodes[parentVal].Right = nodes[childVal]
		}

		// Memo this node became a child
		children[childVal] = true
	}

	// 2. Finding the Root
	for _, desc := range descriptions {
		parentVal := desc[0]
		if !children[parentVal] {
			return nodes[parentVal]
		}
	}

	return nil
}
