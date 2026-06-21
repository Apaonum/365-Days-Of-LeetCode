/*
1840. Maximum Building Height | Hard

You want to build n new buildings in a city. The new buildings will be built in a line and are labeled from 1 to n.

However, there are city restrictions on the heights of the new buildings:

The height of each building must be a non-negative integer.
The height of the first building must be 0.
The height difference between any two adjacent buildings cannot exceed 1.
Additionally, there are city restrictions on the maximum height of specific buildings. These restrictions are given as a 2D integer array restrictions where restrictions[i] = [idi, maxHeighti] indicates that building idi must have a height less than or equal to maxHeighti.

It is guaranteed that each building will appear at most once in restrictions, and building 1 will not be in restrictions.

Return the maximum possible height of the tallest building.

Example 1:
Input: n = 5, restrictions = [[2,1],[4,1]]
Output: 2
Explanation: The green area in the image indicates the maximum allowed height for each building.
We can build the buildings with heights [0,1,2,1,2], and the tallest building has a height of 2.

Example 2:
Input: n = 6, restrictions = []
Output: 5
Explanation: The green area in the image indicates the maximum allowed height for each building.
We can build the buildings with heights [0,1,2,3,4,5], and the tallest building has a height of 5.

Example 3:
Input: n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
Output: 5
Explanation: The green area in the image indicates the maximum allowed height for each building.
We can build the buildings with heights [0,1,2,3,3,4,4,5,4,3], and the tallest building has a height of 5.
 
Constraints:

2 <= n <= 10^9
0 <= restrictions.length <= min(n - 1, 10^5)
2 <= id(i) <= n
id(i) is unique.
0 <= maxHeighti <= 10^9
*/ 

package main

import "sort"

func maxBuilding(n int, restrictions [][]int) int {
	// 1. เพิ่มเสาหลักตึกแรก
	restrictions = append(restrictions, []int{1, 0})
	
	// เรียงลำดับเสาตามหมายเลขตึก
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})
	
	// เพิ่มเสาหลักตึกสุดท้าย (ถ้ายังไม่มี)
	if restrictions[len(restrictions)-1][0] != n {
		restrictions = append(restrictions, []int{n, 1000000000}) // สูงสุดไว้ก่อน
	}

	m := len(restrictions)

	// 2. Forward Pass: กวาดซ้ายไปขวา
	for i := 1; i < m; i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]
		limit := restrictions[i-1][1] + dist
		if limit < restrictions[i][1] {
			restrictions[i][1] = limit
		}
	}

	// 3. Backward Pass: กวาดขวามาซ้าย
	for i := m - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]
		limit := restrictions[i+1][1] + dist
		if limit < restrictions[i][1] {
			restrictions[i][1] = limit
		}
	}

	// 4. หาจุดยอดเขา (Peak) ที่สูงที่สุดระหว่างคู่เสา
	maxHeight := 0
	for i := 0; i < m-1; i++ {
		id1, h1 := restrictions[i][0], restrictions[i][1]
		id2, h2 := restrictions[i+1][0], restrictions[i+1][1]
		
		// สมการหาความสูงยอดเขา
		peak := (h1 + h2 + (id2 - id1)) / 2
		if peak > maxHeight {
			maxHeight = peak
		}
	}

	return maxHeight
}