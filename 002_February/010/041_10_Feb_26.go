/*
3739. Count Subarrays With Majority Element II | Hard
You are given an integer array nums and an integer target.

Return the number of subarrays of nums in which target is the majority element.

The majority element of a subarray is the element that appears strictly more than half of the times in that subarray.

 

Example 1:

Input: nums = [1,2,2,3], target = 2

Output: 5

Explanation:

Valid subarrays with target = 2 as the majority element:

nums[1..1] = [2]
nums[2..2] = [2]
nums[1..2] = [2,2]
nums[0..2] = [1,2,2]
nums[1..3] = [2,2,3]
So there are 5 such subarrays.

Example 2:

Input: nums = [1,1,1,1], target = 1

Output: 10

Explanation:

​​​​​​​All 10 subarrays have 1 as the majority element.

Example 3:

Input: nums = [1,2,3], target = 4

Output: 0

Explanation:

target = 4 does not appear in nums at all. Therefore, there cannot be any subarray where 4 is the majority element. Hence the answer is 0.

 

Constraints:

1 <= nums.length <= 10^​​​​​​​5
1 <= nums[i] <= 10^​​​​​​​9
1 <= target <= 10^9
*/ 

package main

func majoritySubarrayCount(nums []int, target int) int {
	var ans int64 = 0
	n := len(nums)

	// ใช้ Array เก็บความถี่ของผลรวมสะสม (Prefix Sum)
	// เนื่องจากผลรวมอาจติดลบได้ถึง -N เราจึงต้องบวกค่า Offset (N) เสมอ
	freq := make([]int, 2*n+1)
	offset := n
	
	freq[offset] = 1 // เริ่มต้น: ผลรวม 0 มีความถี่ 1 ครั้ง
	
	currentSum := 0
	var countSmaller int64 = 0 // ตัวแปรวิเศษ เก็บจำนวนค่าที่น้อยกว่าปัจจุบัน

	for _, num := range nums {
		if num == target {
			// เดินขึ้น (+1): 
			// ค่าที่น้อยกว่าเรา จะ "เพิ่มขึ้น" เท่ากับจำนวนความถี่ของผลรวมที่เราเพิ่งจากมา
			countSmaller += int64(freq[currentSum+offset])
			currentSum++
		} else {
			// เดินลง (-1):
			// ค่าที่น้อยกว่าเรา จะ "ลดลง" เท่ากับจำนวนความถี่ของผลรวมที่เรากำลังจะไปเหยียบ
			currentSum--
			countSmaller -= int64(freq[currentSum+offset])
		}
		
		ans += countSmaller
		freq[currentSum+offset]++
	}

	return int(ans)
}