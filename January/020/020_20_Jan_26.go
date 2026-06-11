/*
3558. Number of Ways to Assign Edge Weights I | Medium
There is an undirected tree with n nodes labeled from 1 to n, rooted at node 1. The tree is represented by a 2D integer array edges of length n - 1, 
where edges[i] = [u(i), v(i)] indicates that there is an edge between nodes u(i) and v(i).

Initially, all edges have a weight of 0. You must assign each edge a weight of either 1 or 2.

The cost of a path between any two nodes u and v is the total weight of all edges in the path connecting them.

Select any one node x at the maximum depth. Return the number of ways to assign edge weights in the path from node 1 to x such that its total cost is odd.

Since the answer may be large, return it modulo 109 + 7.

Note: Ignore all edges not in the path from node 1 to x.

 

Example 1:

Input: edges = [[1,2]]

Output: 1

Explanation:

The path from Node 1 to Node 2 consists of one edge (1 → 2).
Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
Example 2:



Input: edges = [[1,2],[1,3],[3,4],[3,5]]

Output: 2

Explanation:

The maximum depth is 2, with nodes 4 and 5 at the same depth. Either node can be selected for processing.
For example, the path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4).
Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
 

Constraints:

2 <= n <= 105
edges.length == n - 1
edges[i] == [u(i), v(i)]
1 <= u(i), v(i) <= n
edges represents a valid tree.
*/ 

package main

func numberWays(edges [][]int) int {
	n := len(edges) + 1
	adj := make([][]int, n+1)
	
	// 1. สร้างกราฟ (Adjacency List)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 2. ใช้ BFS หาความลึกสูงสุด (Max Depth)
	maxDepth := 0
	q := []int{1}
	visited := make([]bool, n+1)
	visited[1] = true
	depths := make([]int, n+1)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:] // Pop ตัวแรกออก

		for _, neighbor := range adj[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				depths[neighbor] = depths[curr] + 1
				if depths[neighbor] > maxDepth {
					maxDepth = depths[neighbor]
				}
				q = append(q, neighbor)
			}
		}
	}

	// 3. คำนวณ 2^(maxDepth - 1) % (10^9 + 7)
	const MOD = 1000000007
	return power(2, maxDepth-1, MOD)
}

// ฟังก์ชันยกกำลังความเร็วสูง O(log N)
func power(base, exp, mod int) int {
	res := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}
		exp = exp >> 1 // หาร 2 ปัดเศษลง
		base = (base * base) % mod
	}
	return res
}