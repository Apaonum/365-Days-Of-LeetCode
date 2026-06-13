/*
135. Candy | Hard
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.



Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.


Constraints:

n == ratings.length
1 <= n <= 2 * 10^4
0 <= ratings[i] <= 2 * 10^4
*/

package main

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)

	// 1. Initial condition, Give 1 candy for everyone
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 2. 1st round: From left side to right side
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 3. 2nd round: From right side to left side, Add each of summaries no need to re-loop again
	totalCandies := candies[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		totalCandies += candies[i]
	}

	return totalCandies
}
