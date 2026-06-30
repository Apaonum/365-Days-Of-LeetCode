/*
1358. Number of Substrings Containing All Three Characters | Medium
Given a string s consisting only of characters a, b and c.

Return the number of substrings containing at least one occurrence of all these characters a, b and c.

Example 1:
Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 

Example 2:
Input: s = "aaacb"
Output: 3
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb". 

Example 3:
Input: s = "abc"
Output: 1
 
Constraints:

3 <= s.length <= 5 x 10^4
s only consists of a, b or c characters.
*/ 

package main

func numberOfSubstrings(s string) int {
	last := [3]int{-1, -1, -1}
	ans := 0

	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i

		if last[0] != -1 && last[1] != -1 && last[2] != -1 {
			minIdx := last[0]
			if last[1] < minIdx {
				minIdx = last[1]
			}
			if last[2] < minIdx {
				minIdx = last[2]
			}
			
			ans += minIdx + 1
		}
	}

	return ans
}