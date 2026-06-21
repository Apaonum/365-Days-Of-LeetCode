/*
3614. Process String with Special Operations II | Hard

You are given a string `s` consisting of lowercase English letters and the special characters: `'*'`, `'#'`, and `'%'`.

You are also given an integer `k`.

Build a new string `result` by processing `s` according to the following rules from left to right:

If the letter is a **lowercase** English letter append it to `result`.
A `'*'` **removes** the last character from `result`, if it exists.
A `'#'` **duplicates** the current `result` and **appends** it to itself.
A `'%'` **reverses** the current `result`.

Return the `k`<sup>th</sup> character of the final string `result`. If `k` is out of the bounds of `result`, return `'.'`.

---

### Example 1:

**Input:** `s = "a#b%*", k = 1`
**Output:** `"a"`

**Explanation:**

| i | s[i] | Operation | Current `result` |
| :--- | :--- | :--- | :--- |
| 0 | 'a' | Append 'a' | `"a"` |
| 1 | '#' | Duplicate `result` | `"aa"` |
| 2 | 'b' | Append 'b' | `"aab"` |
| 3 | '%' | Reverse `result` | `"baa"` |
| 4 | '*' | Remove the last character | `"ba"` |

The final `result` is `"ba"`. The character at index `k = 1` is `'a'`.

### Example 2:

**Input:** `s = "cd%#*#", k = 3`
**Output:** `"d"`

**Explanation:**

| i | s[i] | Operation | Current `result` |
| :--- | :--- | :--- | :--- |
| 0 | 'c' | Append 'c' | `"c"` |
| 1 | 'd' | Append 'd' | `"cd"` |
| 2 | '%' | Reverse `result` | `"dc"` |
| 3 | '#' | Duplicate `result` | `"dcdc"` |
| 4 | '*' | Remove the last character | `"dcd"` |
| 5 | '#' | Duplicate `result` | `"dcddcd"` |

The final `result` is `"dcddcd"`. The character at index `k = 3` is `'d'`.

### Example 3:

**Input:** `s = "z*#", k = 0`
**Output:** `"."`

**Explanation:**

| i | s[i] | Operation | Current `result` |
| :--- | :--- | :--- | :--- |
| 0 | 'z' | Append 'z' | `"z"` |
| 1 | '*' | Remove the last character | `""` |
| 2 | '#' | Duplicate the string | `""` |

The final `result` is `""`. Since index `k = 0` is out of bounds, the output is `'.'`.

---

### Constraints:

`1 <= s.length <= 10^5`
`s` consists of only lowercase English letters and special characters `'*'`, `'#'`, and `'%'`.
`0 <= k <= 10^15`
The length of `result` after processing `s` will not exceed `10^15`.
*/

package main

func kthCharacter(s string, k int64) byte {
	n := len(s)
	lengths := make([]int64, n)
	var currLen int64 = 0

	// 1. เดินหน้าจดความยาวในแต่ละสเต็ป
	for i := 0; i < n; i++ {
		char := s[i]
		if char >= 'a' && char <= 'z' {
			currLen++
		} else if char == '*' {
			if currLen > 0 {
				currLen--
			}
		} else if char == '#' {
			currLen *= 2
		} else if char == '%' {
			// ความยาวเท่าเดิม
		}
		lengths[i] = currLen
	}

	// ถ้า k อยู่นอกขอบเขตของความยาวสุดท้าย
	if k >= lengths[n-1] {
		return '.'
	}

	// 2. เดินถอยหลังเพื่อย้อนรอยหาต้นตอ
	for i := n - 1; i >= 0; i-- {
		char := s[i]

		// หาความยาวก่อนหน้าที่จะทำคำสั่งนี้
		var prevLen int64 = 0
		if i > 0 {
			prevLen = lengths[i-1]
		}

		if char >= 'a' && char <= 'z' {
			// ถ้า k ชี้มาที่ตำแหน่งที่เพิ่งถูกเพิ่มเข้ามาพอดี เจอตัวแล้ว!
			if k == prevLen {
				return char
			}
		} else if char == '*' {
			// k ไม่ได้รับผลกระทบจากการลบตัวท้าย
		} else if char == '#' {
			if prevLen > 0 {
				k %= prevLen // ตบ k กลับมาฝั่งต้นฉบับ
			}
		} else if char == '%' {
			if prevLen > 0 {
				k = prevLen - 1 - k // สลับตำแหน่งหน้าหลังคืน
			}
		}
	}

	return '.'
}
