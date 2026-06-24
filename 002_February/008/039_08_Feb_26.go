/*
3700. Number of ZigZag Arrays II | Hard
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
[5, 4, 5]

Example 2:

Input: n = 3, l = 1, r = 3

Output: 10

Explanation:

​​​​​​​There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:

[1, 2, 1], [1, 3, 1], [1, 3, 2]
[2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
[3, 1, 2], [3, 1, 3], [3, 2, 3]
All arrays meet the ZigZag conditions.

Constraints:

3 <= n <= 10^9
1 <= l < r <= 75
*/ 

package main

// ฟังก์ชันคูณเมทริกซ์ (Matrix Multiplication)
func multiplyMatrix(A, B [][]int) [][]int {
	size := len(A)
	C := make([][]int, size)
	for i := 0; i < size; i++ {
		C[i] = make([]int, size)
		for k := 0; k < size; k++ {
			// ข้ามการคำนวณถ้าเป็น 0 เพื่อรีด Performance
			if A[i][k] == 0 {
				continue
			}
			for j := 0; j < size; j++ {
				// ใช้ int64 ป้องกันการ Overflow ก่อนทำ Modulo
				C[i][j] = int((int64(C[i][j]) + int64(A[i][k])*int64(B[k][j])) % 1000000007)
			}
		}
	}
	return C
}

// ฟังก์ชันยกกำลังเมทริกซ์ (Matrix Fast Exponentiation)
func powerMatrix(A [][]int, p int) [][]int {
	size := len(A)
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		res[i][i] = 1 // Identity Matrix
	}
	base := A
	for p > 0 {
		if p%2 == 1 {
			res = multiplyMatrix(res, base)
		}
		base = multiplyMatrix(base, base)
		p /= 2
	}
	return res
}

func zigzagArrays(n int, l int, r int) int {
	K := r - l + 1
	if K < 2 {
		return 0
	}

	size := 2 * K
	M := make([][]int, size)
	for i := 0; i < size; i++ {
		M[i] = make([]int, size)
	}

	// สร้างสมการการก้าว (Transition Matrix)
	// แถวที่ 0 ถึง K-1 คือฝั่ง "เดินขึ้น" (ขึ้นอยู่กับฝั่ง "เดินลง" ในสเต็ปที่แล้ว)
	for row := 0; row < K; row++ {
		for col := K; col < K+row; col++ {
			M[row][col] = 1
		}
	}

	// แถวที่ K ถึง 2K-1 คือฝั่ง "เดินลง" (ขึ้นอยู่กับฝั่ง "เดินขึ้น" ในสเต็ปที่แล้ว)
	for row := K; row < 2*K; row++ {
		for col := (row - K) + 1; col < K; col++ {
			M[row][col] = 1
		}
	}

	// นำเมทริกซ์ไปยกกำลัง n-2 แบบสายฟ้าแลบ
	M_pow := powerMatrix(M, n-2)

	// สร้างสถานะเริ่มต้นสำหรับอาร์เรย์ความยาว 2
	V2 := make([]int, size)
	for j := 1; j <= K; j++ {
		V2[j-1] = j - 1      // เดินขึ้นมาหา j มีวิธี j-1 แบบ
		V2[K+j-1] = K - j    // เดินลงมาหา j มีวิธี K-j แบบ
	}

	ans := 0
	const MOD = 1000000007
	
	// คูณเมทริกซ์ผลลัพธ์เข้ากับสถานะเริ่มต้น V2
	for i := 0; i < size; i++ {
		val := 0
		for j := 0; j < size; j++ {
			val = int((int64(val) + int64(M_pow[i][j])*int64(V2[j])) % MOD)
		}
		ans = (ans + val) % MOD
	}

	return ans
}