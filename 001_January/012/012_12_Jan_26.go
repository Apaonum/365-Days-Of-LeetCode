/*
2574. Left and Right Sum Differences | Easy
You are given a 0-indexed integer array nums of size n.

Define two arrays leftSum and rightSum where:

leftSum[i] is the sum of elements to the left of the index i in the array nums. If there is no such element, leftSum[i] = 0.
rightSum[i] is the sum of elements to the right of the index i in the array nums. If there is no such element, rightSum[i] = 0.
Return an integer array answer of size n where answer[i] = |leftSum[i] - rightSum[i]|.

Example 1:

Input: nums = [10,4,8,3]
Output: [15,1,11,22]
Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].
Example 2:

Input: nums = [1]
Output: [0]
Explanation: The array leftSum is [0] and the array rightSum is [0].
The array answer is [|0 - 0|] = [0].
 

Constraints:

1 <= nums.length <= 1000
1 <= nums[i] <= 10^5
*/ 

package main

func leftRightDifference(nums []int) []int {
    // 1. Total Summary
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    
    leftSum := 0
    ans := make([]int, len(nums))
    
    // 2. Check 
    for i, num := range nums {
        rightSum := totalSum - leftSum - num
        
        diff := leftSum - rightSum
        if diff < 0 {
            diff = -diff
        }
        ans[i] = diff
        
        leftSum += num
    }
    
    return ans
}