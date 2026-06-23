/*
3699. Number of ZigZag Arrays I | Hard
You are given three integers n, l, and r.

A ZigZag array of length n is defined as follows:

Each element lies in the range [l, r].
No two adjacent elements are equal.
No three consecutive elements form a strictly increasing or strictly decreasing sequence.
Return the total number of valid ZigZag arrays.

Since the answer may be large, return it modulo 10^9 + 7.

A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

A sequence is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

 

Example 1:
Input: n = 3, l = 4, r = 5
Output: 2
Explanation:
There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:
[4, 5, 4]
[5, 4, 5]​​​​​​​

Example 2:
Input: n = 3, l = 1, r = 3
Output: 10
Explanation:
There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:
[1, 2, 1], [1, 3, 1], [1, 3, 2]
[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
[3, 1, 2], [3, 1, 3], [3, 2, 3]
All arrays meet the ZigZag conditions.

Constraints:

3 <= n <= 2000
1 <= l < r <= 2000
*/ 

package main

func zigzagArrays(n int, l int, r int) int {
	K := r - l + 1
	// ถ้ามีตัวเลขให้เลือกไม่ถึง 2 ตัว ก็ทำ ZigZag ไม่ได้แน่นอน
	if K < 2 {
		return 0
	}
	
	const MOD = 1000000007

	// เก็บสถานะของรอบที่แล้ว
	prevUp := make([]int, K+2)
	prevDown := make([]int, K+2)

	// Base Case: สำหรับอาร์เรย์ความยาว 2
	for j := 1; j <= K; j++ {
		prevUp[j] = j - 1    // เดินขึ้นมาหา j (ตัวหน้าต้องน้อยกว่า j ซึ่งมี j-1 ตัว)
		prevDown[j] = K - j  // เดินลงมาหา j (ตัวหน้าต้องมากกว่า j ซึ่งมี K-j ตัว)
	}

	// เริ่มสร้างอาร์เรย์ความยาว 3 ไปจนถึง n
	for i := 3; i <= n; i++ {
		currUp := make([]int, K+2)
		currDown := make([]int, K+2)

		// 1. คำนวณฝั่ง "เดินขึ้น" ด้วยผลรวมสะสมจากซ้ายไปขวา (Prefix Sum ของ prevDown)
		sum := 0
		for j := 1; j <= K; j++ {
			currUp[j] = sum
			sum = (sum + prevDown[j]) % MOD
		}

		// 2. คำนวณฝั่ง "เดินลง" ด้วยผลรวมสะสมจากขวามาซ้าย (Suffix Sum ของ prevUp)
		sum = 0
		for j := K; j >= 1; j-- {
			currDown[j] = sum
			sum = (sum + prevUp[j]) % MOD
		}

		// อัปเดตสถานะเพื่อใช้ในรอบถัดไป
		prevUp = currUp
		prevDown = currDown
	}

	// นับผลรวมวิธีทั้งหมดในตอนจบ
	ans := 0
	for j := 1; j <= K; j++ {
		ans = (ans + prevUp[j]) % MOD
		ans = (ans + prevDown[j]) % MOD
	}
	
	return ans
}