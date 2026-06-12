/*
3559. Number of Ways to Assign Edge Weights II | Hard
There is an undirected tree with n nodes labeled from 1 to n, rooted at node 1. The tree is represented by a 2D integer array edges of length n - 1, 
where edges[i] = [u(i), v(i)] indicates that there is an edge between nodes u(i) and v(i).

Initially, all edges have a weight of 0. You must assign each edge a weight of either 1 or 2.

The cost of a path between any two nodes u and v is the total weight of all edges in the path connecting them.

You are given a 2D integer array queries. For each queries[i] = [u(i), v(i)], 
determine the number of ways to assign weights to edges in the path such that the cost of the path between u(i) and v(i) is odd.

Return an array answer, where answer[i] is the number of valid assignments for queries[i].

Since the answer may be large, apply modulo 109 + 7 to each answer[i].

Note: For each query, disregard all edges not in the path between node u(i) and v(i).

 

Example 1:

Input: edges = [[1,2]], queries = [[1,1],[1,2]]

Output: [0,1]

Explanation:

Query [1,1]: The path from Node 1 to itself consists of no edges, so the cost is 0. Thus, the number of valid assignments is 0.
Query [1,2]: The path from Node 1 to Node 2 consists of one edge (1 → 2). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.

Example 2:

Input: edges = [[1,2],[1,3],[3,4],[3,5]], queries = [[1,4],[3,4],[2,5]]

Output: [2,1,4]

Explanation:

Query [1,4]: The path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4). Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
Query [3,4]: The path from Node 3 to Node 4 consists of one edge (3 → 4). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
Query [2,5]: The path from Node 2 to Node 5 consists of three edges (2 → 1, 1 → 3, and 3 → 5). Assigning (1,2,2), (2,1,2), (2,2,1), or (1,1,1) makes the cost odd. Thus, the number of valid assignments is 4.
 

Constraints:

2 <= n <= 10^5
edges.length == n - 1
edges[i] == [u(i), v(i)]
1 <= queries.length <= 10^5
queries[i] == [u(i), v(i)]
1 <= u(i), v(i) <= n
edges represents a valid tree.
*/ 

package main

func answerQueries(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	const MOD = 1000000007
	const LOG = 19 

	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}

	depth := make([]int, n+1)
	up := make([][]int, n+1)
	for i := range up {
		up[i] = make([]int, LOG)
	}

	q := []int{1}
	visited := make([]bool, n+1)
	visited[1] = true
	depth[1] = 0
	up[1][0] = 1

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, neighbor := range adj[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				depth[neighbor] = depth[curr] + 1
				up[neighbor][0] = curr
				q = append(q, neighbor)
			}
		}
	}

	for j := 1; j < LOG; j++ {
		for i := 1; i <= n; i++ {
			up[i][j] = up[up[i][j-1]][j-1]
		}
	}


	getLCA := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		
		diff := depth[u] - depth[v]
		for j := 0; j < LOG; j++ {
			if (diff & (1 << j)) != 0 {
				u = up[u][j]
			}
		}
		
		if u == v {
			return u
		}

		for j := LOG - 1; j >= 0; j-- {
			if up[u][j] != up[v][j] {
				u = up[u][j]
				v = up[v][j]
			}
		}
		
		return up[u][0]
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		if u == v {
			ans[i] = 0
			continue
		}

		lca := getLCA(u, v)
		distance := depth[u] + depth[v] - 2*depth[lca]

		if distance == 0 {
			ans[i] = 0
		} else {
			ans[i] = pow2[distance-1]
		}
	}

	return ans
}