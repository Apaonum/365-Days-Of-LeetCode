/*
3635. Earliest Finish Time for Land and Water Rides II | Medium
You are given two categories of theme park attractions: land rides and water rides.

Land rides
landStartTime[i] – the earliest time the ith land ride can be boarded.
landDuration[i] – how long the ith land ride lasts.
Water rides
waterStartTime[j] – the earliest time the jth water ride can be boarded.
waterDuration[j] – how long the jth water ride lasts.
A tourist must experience exactly one ride from each category, in either order.

A ride may be started at its opening time or any later moment.
If a ride is started at time t, it finishes at time t + duration.
Immediately after finishing one ride the tourist may board the other (if it is already open) or wait until it opens.
Return the earliest possible time at which the tourist can finish both rides.



Example 1:

Input: landStartTime = [2,8], landDuration = [4,1], waterStartTime = [6], waterDuration = [3]

Output: 9

Explanation:​​​​​​​

Plan A (land ride 0 → water ride 0):
Start land ride 0 at time landStartTime[0] = 2. Finish at 2 + landDuration[0] = 6.
Water ride 0 opens at time waterStartTime[0] = 6. Start immediately at 6, finish at 6 + waterDuration[0] = 9.
Plan B (water ride 0 → land ride 1):
Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
Land ride 1 opens at landStartTime[1] = 8. Start at time 9, finish at 9 + landDuration[1] = 10.
Plan C (land ride 1 → water ride 0):
Start land ride 1 at time landStartTime[1] = 8. Finish at 8 + landDuration[1] = 9.
Water ride 0 opened at waterStartTime[0] = 6. Start at time 9, finish at 9 + waterDuration[0] = 12.
Plan D (water ride 0 → land ride 0):
Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
Land ride 0 opened at landStartTime[0] = 2. Start at time 9, finish at 9 + landDuration[0] = 13.
Plan A gives the earliest finish time of 9.

Example 2:

Input: landStartTime = [5], landDuration = [3], waterStartTime = [1], waterDuration = [10]

Output: 14

Explanation:​​​​​​​

Plan A (water ride 0 → land ride 0):
Start water ride 0 at time waterStartTime[0] = 1. Finish at 1 + waterDuration[0] = 11.
Land ride 0 opened at landStartTime[0] = 5. Start immediately at 11 and finish at 11 + landDuration[0] = 14.
Plan B (land ride 0 → water ride 0):
Start land ride 0 at time landStartTime[0] = 5. Finish at 5 + landDuration[0] = 8.
Water ride 0 opened at waterStartTime[0] = 1. Start immediately at 8 and finish at 8 + waterDuration[0] = 18.
Plan A provides the earliest finish time of 14.​​​​​​​



Constraints:

1 <= n, m <= 5 * 10^4
landStartTime.length == landDuration.length == n
waterStartTime.length == waterDuration.length == m
1 <= landStartTime[i], landDuration[i], waterStartTime[j], waterDuration[j] <= 10^5
*/

// Solution: Have 2 section -> 1. water to land time and land to water time

package main

import (
	"math"
	"sort"
)

type Ride struct {
	s int
	d int
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// Check minimum time of land to water or water to land
	ans1 := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	ans2 := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(ans1, ans2)
}

func solve(firstS, firstD, secondS, secondD []int) int {
	n := len(secondS)
	second := make([]Ride, n)
	for i := 0; i < n; i++ {
		second[i] = Ride{s: secondS[i], d: secondD[i]}
	}

	// Sorted the second ride by start time
	sort.Slice(second, func(i, j int) bool {
		return second[i].s < second[j].s
	})

	// Defined prefix arrays and find minimum duration (from left to right)
	prefMinD := make([]int, n)
	prefMinD[0] = second[0].d
	for i := 1; i < n; i++ {
		prefMinD[i] = min(prefMinD[i-1], second[i].d)
	}

	// Defined suffix arrays and find minimum finish time (s + d)
	suffMinEnd := make([]int, n)
	suffMinEnd[n-1] = second[n-1].s + second[n-1].d
	for i := n - 2; i >= 0; i-- {
		suffMinEnd[i] = min(suffMinEnd[i+1], second[i].s+second[i].d)
	}

	best := math.MaxInt64

	// loop
	for i := 0; i < len(firstS); i++ {
		T := firstS[i] + firstD[i]

		idx := sort.Search(n, func(k int) bool {
			return second[k].s >= T
		})

		ans := math.MaxInt64

		if idx < n {
			ans = min(ans, suffMinEnd[idx])
		}

		if idx > 0 {
			ans = min(ans, T+prefMinD[idx-1])
		}

		best = min(best, ans)
	}

	return best
}
