/*
3691. Maximum Total Subarray Value II | Hard
You are given an integer array nums of length n and an integer k.

You must select exactly k distinct subarrays nums[l..r] of nums. Subarrays may overlap, but the exact same subarray (same l and r) cannot be chosen more than once.

The value of a subarray nums[l..r] is defined as: max(nums[l..r]) - min(nums[l..r]).

The total value is the sum of the values of all chosen subarrays.

Return the maximum possible total value you can achieve.



Example 1:

Input: nums = [1,3,2], k = 2

Output: 4

Explanation:

One optimal approach is:

Choose nums[0..1] = [1, 3]. The maximum is 3 and the minimum is 1, giving a value of 3 - 1 = 2.
Choose nums[0..2] = [1, 3, 2]. The maximum is still 3 and the minimum is still 1, so the value is also 3 - 1 = 2.
Adding these gives 2 + 2 = 4.

Example 2:

Input: nums = [4,2,5,1], k = 3

Output: 12

Explanation:

One optimal approach is:

Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, giving a value of 5 - 1 = 4.
Choose nums[1..3] = [2, 5, 1]. The maximum is 5 and the minimum is 1, so the value is also 4.
Choose nums[2..3] = [5, 1]. The maximum is 5 and the minimum is 1, so the value is again 4.
Adding these gives 4 + 4 + 4 = 12.


Constraints:

1 <= n == nums.length <= 5 * 10​​​​​​​^4
0 <= nums[i] <= 10^9
1 <= k <= min(10^5, n * (n + 1) / 2)
*/

package main

// โครงสร้าง Fenwick Tree ที่รองรับการอัปเดตเป็นช่วง (Range Update) และหาผลรวมเป็นช่วง (Range Sum)
type RangeBIT struct {
	bit1 []int64
	bit2 []int64
	n    int
}

func NewRangeBIT(n int) *RangeBIT {
	return &RangeBIT{bit1: make([]int64, n+2), bit2: make([]int64, n+2), n: n}
}

func (b *RangeBIT) addPoint(idx int, val int64) {
	val2 := val * int64(idx-1)
	for i := idx; i <= b.n; i += i & -i {
		b.bit1[i] += val
		b.bit2[i] += val2
	}
}

func (b *RangeBIT) AddRange(l, r int, val int64) {
	b.addPoint(l, val)
	b.addPoint(r+1, -val)
}

func (b *RangeBIT) queryPrefix(idx int) int64 {
	var sum1, sum2 int64
	for i := idx; i > 0; i -= i & -i {
		sum1 += b.bit1[i]
		sum2 += b.bit2[i]
	}
	return sum1*int64(idx) - sum2
}

func maxSubarrayValue(nums []int, k int) int64 {
	n := len(nums)

	// ฟังก์ชันนับจำนวน Subarray ที่มีผลต่าง (Max - Min) >= v
	countGreaterOrEqual := func(v int) int64 {
		var ans int64 = 0
		maxDq := make([]int, 0, n)
		minDq := make([]int, 0, n)
		l := 0
		for r := 0; r < n; r++ {
			for len(maxDq) > 0 && nums[maxDq[len(maxDq)-1]] <= nums[r] {
				maxDq = maxDq[:len(maxDq)-1]
			}
			maxDq = append(maxDq, r)
			for len(minDq) > 0 && nums[minDq[len(minDq)-1]] >= nums[r] {
				minDq = minDq[:len(minDq)-1]
			}
			minDq = append(minDq, r)

			// หด Window ฝั่งซ้าย ถ้าระยะห่าง >= v
			for len(maxDq) > 0 && len(minDq) > 0 && nums[maxDq[0]]-nums[minDq[0]] >= v {
				if maxDq[0] == l {
					maxDq = maxDq[1:]
				}
				if minDq[0] == l {
					minDq = minDq[1:]
				}
				l++
			}
			// จำนวน Subarray ที่ผลต่าง < v
			ans += int64(r - l + 1)
		}
		total := int64(n) * int64(n+1) / 2
		return total - ans
	}

	// 1. Binary Search หาเกณฑ์ V
	low, high, bestV := 0, 1000000000, 0
	for low <= high {
		mid := low + (high-low)/2
		if countGreaterOrEqual(mid) >= int64(k) {
			bestV = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// 2. ใช้ BIT และ Monotonic Stack หาผลรวมของอันที่ได้คะแนน > bestV (vStrict)
	vStrict := bestV + 1
	maxBit := NewRangeBIT(n)
	minBit := NewRangeBIT(n)

	maxStack := make([]int, 0, n)
	minStack := make([]int, 0, n)
	maxDq := make([]int, 0, n)
	minDq := make([]int, 0, n)

	var totalSum int64 = 0
	var validCount int64 = 0
	l := 1

	for r := 1; r <= n; r++ {
		val := int64(nums[r-1])

		// อัปเดต Max BIT
		for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]-1] <= nums[r-1] {
			popped := maxStack[len(maxStack)-1]
			maxStack = maxStack[:len(maxStack)-1]
			prev := 0
			if len(maxStack) > 0 {
				prev = maxStack[len(maxStack)-1]
			}
			maxBit.AddRange(prev+1, popped, val-int64(nums[popped-1]))
		}
		maxStack = append(maxStack, r)
		maxBit.AddRange(r, r, val)

		// อัปเดต Min BIT
		for len(minStack) > 0 && nums[minStack[len(minStack)-1]-1] >= nums[r-1] {
			popped := minStack[len(minStack)-1]
			minStack = minStack[:len(minStack)-1]
			prev := 0
			if len(minStack) > 0 {
				prev = minStack[len(minStack)-1]
			}
			minBit.AddRange(prev+1, popped, val-int64(nums[popped-1]))
		}
		minStack = append(minStack, r)
		minBit.AddRange(r, r, val)

		// คุมหน้าต่าง Sliding Window สำหรับ vStrict
		for len(maxDq) > 0 && nums[maxDq[len(maxDq)-1]-1] <= nums[r-1] {
			maxDq = maxDq[:len(maxDq)-1]
		}
		maxDq = append(maxDq, r)

		for len(minDq) > 0 && nums[minDq[len(minDq)-1]-1] >= nums[r-1] {
			minDq = minDq[:len(minDq)-1]
		}
		minDq = append(minDq, r)

		for len(maxDq) > 0 && len(minDq) > 0 && nums[maxDq[0]-1]-nums[minDq[0]-1] >= vStrict {
			if maxDq[0] == l {
				maxDq = maxDq[1:]
			}
			if minDq[0] == l {
				minDq = minDq[1:]
			}
			l++
		}

		validCount += int64(l - 1)
		totalSum += maxBit.queryPrefix(l-1) - minBit.queryPrefix(l-1)
	}

	// คืนค่า: ผลรวมของอันที่ > V + (เติมอันที่ = V ให้ครบจำนวน k)
	return totalSum + (int64(k)-validCount)*int64(bestV)
}
