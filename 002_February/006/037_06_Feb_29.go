/*
1189. Maximum Number of Balloons | Easy
Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

Example 1:
Input: text = "nlaebolko"
Output: 1

Example 2:
Input: text = "loonbalxballpoon"
Output: 2

Example 3:
Input: text = "leetcode"
Output: 0
 

Constraints:

1 <= text.length <= 10^4
text consists of lower case English letters only.
*/ 

package main

func maxNumberOfBalloons(text string) int {
	counts := make([]int, 26)
	for _, char := range text {
		counts[char-'a']++
	}

	b := counts['b'-'a']
	a := counts['a'-'a']
	l := counts['l'-'a'] / 2 
	o := counts['o'-'a'] / 2 
	n := counts['n'-'a']

	return min(b, min(a, min(l, min(o, n))))
}