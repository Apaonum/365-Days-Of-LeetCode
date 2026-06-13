package main

import (
	"testing"
)

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		descriptions [][]int
		expected     *TreeNode
	}{
		{
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			// 50 -> left:20, right:80 | 20 -> left:15, right:17 | 80 -> left:19
			expected: &TreeNode{Val: 50,
				Left: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 17},
				},
				Right: &TreeNode{Val: 80,
					Left: &TreeNode{Val: 19},
				},
			},
		},
		{
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			// 1 -> left:2 | 2 -> right:3 | 3 -> left:4
			expected: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		result := createBinaryTree(tt.descriptions)
		if !isSameTree(result, tt.expected) {
			t.Errorf("createBinaryTree(%v) tree mismatch", tt.descriptions)
		}
	}
}
