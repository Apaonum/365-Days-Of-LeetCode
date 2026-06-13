/*
3161. Block Placement Queries | Hard

There exists an infinite number line, with its origin at 0 and extending towards the positive x-axis.

You are given a 2D array queries, which contains two types of queries:

For a query of type 1, queries[i] = [1, x]. Build an obstacle at distance x from the origin. It is guaranteed that there is no obstacle at distance x when the query is asked.

For a query of type 2, queries[i] = [2, x, sz]. Check if it is possible to place a block of size sz anywhere in the range [0, x] on the line, such that the block entirely lies in the range [0, x]. A block cannot be placed if it intersects with any obstacle, but it may touch it. Note that you do not actually place the block. Queries are separate.
Return a boolean array results, where results[i] is true if you can place the block specified in the ith query of type 2, and false otherwise. 

Example 1:
Input: queries = [[1,2],[2,3,3],[2,3,1],[2,2,2]]
Output: [false,true,true]

Explanation:
For query 0, place an obstacle at x = 2. A block of size at most 2 can be placed before x = 3.

Example 2:
Input: queries = [[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]
Output: [true,true,false]

Explanation:
Place an obstacle at x = 7 for query 0. A block of size at most 7 can be placed before x = 7.
Place an obstacle at x = 2 for query 2. Now, a block of size at most 5 can be placed before x = 7, and a block of size at most 2 before x = 2.
 

Constraints:

1 <= queries.length <= 15 * 10^4
2 <= queries[i].length <= 3
1 <= queries[i][0] <= 2
1 <= x, sz <= min(5 * 10^4, 3 * queries.length)
The input is generated such that for queries of type 1, no obstacle exists at distance x when the query is asked.
The input is generated such that there is at least one query of type 2.
*/ 

package main

const INF = 1_000_000_000
type Node struct {
	minObs int
	maxObs int
	maxGap int
}

func getResults(queries [][]int) []bool {
	// 1. Created fit Segment Tree
	maxX := 0
	for _, q := range queries {
		if q[1] > maxX {
			maxX = q[1]
		}
	}

	// Created 4x tree of area
	tree := make([]Node, 4*(maxX+1))

	// Tree build func
	var build func(node, L, R int)
	build = func(node, L, R int) {
		if L == R {
			if L == 0 {
				tree[node] = Node{0, 0, 0}
			} else {
				tree[node] = Node{INF, -INF, 0}
			}
			return
		}
		mid := (L + R) / 2
		build(2*node, L, mid)
		build(2*node+1, mid+1, R)

		lc, rc := 2*node, 2*node+1
		tree[node].minObs = min(tree[lc].minObs, tree[rc].minObs)
		tree[node].maxObs = max(tree[lc].maxObs, tree[rc].maxObs)
		tree[node].maxGap = max(tree[lc].maxGap, tree[rc].maxGap)
		if tree[lc].maxObs != -INF && tree[rc].minObs != INF {
			gap := tree[rc].minObs - tree[lc].maxObs
			tree[node].maxGap = max(tree[node].maxGap, gap)
		}
	}

	build(1, 0, maxX)

	// Update func
	var update func(node, L, R, pos int)
	update = func(node, L, R, pos int) {
		if L == R {
			tree[node] = Node{pos, pos, 0}
			return
		}
		mid := (L + R) / 2
		if pos <= mid {
			update(2*node, L, mid, pos)
		} else {
			update(2*node+1, mid+1, R, pos)
		}
		lc, rc := 2*node, 2*node+1
		tree[node].minObs = min(tree[lc].minObs, tree[rc].minObs)
		tree[node].maxObs = max(tree[lc].maxObs, tree[rc].maxObs)
		tree[node].maxGap = max(tree[lc].maxGap, tree[rc].maxGap)
		if tree[lc].maxObs != -INF && tree[rc].minObs != INF {
			gap := tree[rc].minObs - tree[lc].maxObs
			tree[node].maxGap = max(tree[node].maxGap, gap)
		}
	}

	// Find biggest gap
	var query func(node, L, R, QL, QR int) Node
	query = func(node, L, R, QL, QR int) Node {
		if QL <= L && R <= QR {
			return tree[node]
		}
		mid := (L + R) / 2
		if QR <= mid {
			return query(2*node, L, mid, QL, QR)
		}
		if QL > mid {
			return query(2*node+1, mid+1, R, QL, QR)
		}
		leftRes := query(2*node, L, mid, QL, QR)
		rightRes := query(2*node+1, mid+1, R, QL, QR)

		res := Node{}
		res.minObs = min(leftRes.minObs, rightRes.minObs)
		res.maxObs = max(leftRes.maxObs, rightRes.maxObs)
		res.maxGap = max(leftRes.maxGap, rightRes.maxGap)
		if leftRes.maxObs != -INF && rightRes.minObs != INF {
			gap := rightRes.minObs - leftRes.maxObs
			res.maxGap = max(res.maxGap, gap)
		}
		return res
	}

	var results []bool
	for _, q := range queries {
		if q[0] == 1 {
			update(1, 0, maxX, q[1])
		} else {
			x, sz := q[1], q[2]
			res := query(1, 0, maxX, 0, x)
			maxPossible := res.maxGap
			if x-res.maxObs > maxPossible {
				maxPossible = x - res.maxObs
			}
			results = append(results, maxPossible >= sz)
		}
	}
	return results
}