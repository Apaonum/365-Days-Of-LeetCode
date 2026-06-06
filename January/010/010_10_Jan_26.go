/*
3753. Total Waviness of Numbers in Range II | Hard
You are given two integers num1 and num2 representing an inclusive range [num1, num2].

The waviness of a number is defined as the total count of its peaks and valleys:

A digit is a peak if it is strictly greater than both of its immediate neighbors.
A digit is a valley if it is strictly less than both of its immediate neighbors.
The first and last digits of a number cannot be peaks or valleys.
Any number with fewer than 3 digits has a waviness of 0.
Return the total sum of waviness for all numbers in the range [num1, num2].


Example 1:

Input: num1 = 120, num2 = 130

Output: 3

Explanation:

In the range [120, 130]:

120: middle digit 2 is a peak, waviness = 1.
121: middle digit 2 is a peak, waviness = 1.
130: middle digit 3 is a peak, waviness = 1.
All other numbers in the range have a waviness of 0.
Thus, total waviness is 1 + 1 + 1 = 3.

Example 2:

Input: num1 = 198, num2 = 202

Output: 3

Explanation:

In the range [198, 202]:

198: middle digit 9 is a peak, waviness = 1.
201: middle digit 0 is a valley, waviness = 1.
202: middle digit 0 is a valley, waviness = 1.
All other numbers in the range have a waviness of 0.
Thus, total waviness is 1 + 1 + 1 = 3.

Example 3:

Input: num1 = 4848, num2 = 4848

Output: 2

Explanation:

Number 4848: the second digit 8 is a peak, and the third digit 4 is a valley, giving a waviness of 2.



Constraints:

1 <= num1 <= num2 <= 10^15
*/

package main

import "strconv"

// DP structure
type Result struct {
	count int64
	sum   int64
}

func totalWaviness(num1 int, num2 int) int {
	// find summary from 0 to num2 minus with 0 to num1-1
	return int(solve(num2) - solve(num1-1))
}

func solve(num int) int64 {
	if num < 100 {
		return 0 // If not bigger than 3 digits will skip. Cause it's not have a peak
	}
	s := strconv.Itoa(num)
	n := len(s)

	// Create an Array 1 dimension for memorize DP (Compatible for maximum State aroud 8,000 blocks)
	visited := make([]bool, 20000)
	memo := make([]Result, 20000)

	var dfs func(idx, p2, p1, isLim, isLead int) Result
	dfs = func(idx, p2, p1, isLim, isLead int) Result {
		if idx == n {
			return Result{1, 0} // finished 1 route (Count 1 number, peak 0)
		}

		// Convert fomulae State 5 index to 1 index (State Compression)
		key := ((((idx*11)+(p2+1))*11+(p1+1))*2+isLim)*2 + isLead
		if visited[key] {
			return memo[key]
		}

		limit := 9
		if isLim == 1 {
			limit = int(s[idx] - '0')
		}

		var totalCount int64 = 0
		var totalSum int64 = 0

		// Add 0 to limit
		for d := 0; d <= limit; d++ {
			nextLim := 0
			if isLim == 1 && d == limit {
				nextLim = 1
			}

			nextLead := 0
			if isLead == 1 && d == 0 {
				nextLead = 1
			}

			nextP2 := p1
			nextP1 := d
			if nextLead == 1 {
				nextP2, nextP1 = -1, -1 // 0 leading not number
			} else if isLead == 1 {
				nextP2, nextP1 = -1, d // out of 0 leading, starting count
			}

			res := dfs(idx+1, nextP2, nextP1, nextLim, nextLead)
			totalCount += res.count

			// Peak or Valley Checked? (Have 3 digits and not leading by 0)
			var wave int64 = 0
			if p2 != -1 && p1 != -1 && nextLead == 0 {
				isPeak := p1 > p2 && p1 > d
				isValley := p1 < p2 && p1 < d
				if isPeak || isValley {
					wave = 1
				}
			}

			// Summary : res + (res count ×  wave)
			totalSum += res.sum + (res.count * wave)
		}

		ans := Result{totalCount, totalSum}
		visited[key] = true
		memo[key] = ans
		return ans
	}

	return dfs(0, -1, -1, 1, 1).sum
}
