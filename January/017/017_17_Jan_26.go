/*
3093. Longest Common Suffix Queries | Hard
You are given two arrays of strings wordsContainer and wordsQuery.

For each wordsQuery[i], you need to find a string from wordsContainer that has the longest common suffix with wordsQuery[i]. If there are two or more strings in wordsContainer that share the longest common suffix, find the string that is the smallest in length. If there are two or more such strings that have the same smallest length, find the one that occurred earlier in wordsContainer.

Return an array of integers ans, where ans[i] is the index of the string in wordsContainer that has the longest common suffix with wordsQuery[i].

 

Example 1:

Input: wordsContainer = ["abcd","bcd","xbcd"], wordsQuery = ["cd","bcd","xyz"]

Output: [1,1,1]

Explanation:

Let's look at each wordsQuery[i] separately:

For wordsQuery[0] = "cd", strings from wordsContainer that share the longest common suffix "cd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
For wordsQuery[1] = "bcd", strings from wordsContainer that share the longest common suffix "bcd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
For wordsQuery[2] = "xyz", there is no string from wordsContainer that shares a common suffix. Hence the longest common suffix is "", that is shared with strings at index 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
Example 2:

Input: wordsContainer = ["abcdefgh","poiuygh","ghghgh"], wordsQuery = ["gh","acbfgh","acbfegh"]

Output: [2,0,2]

Explanation:

Let's look at each wordsQuery[i] separately:

For wordsQuery[0] = "gh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.
For wordsQuery[1] = "acbfgh", only the string at index 0 shares the longest common suffix "fgh". Hence it is the answer, even though the string at index 2 is shorter.
For wordsQuery[2] = "acbfegh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.
 

Constraints:

1 <= wordsContainer.length, wordsQuery.length <= 10^4
1 <= wordsContainer[i].length <= 5 * 10^3
1 <= wordsQuery[i].length <= 5 * 10^3
wordsContainer[i] consists only of lowercase English letters.
wordsQuery[i] consists only of lowercase English letters.
Sum of wordsContainer[i].length is at most 5 * 10^5.
Sum of wordsQuery[i].length is at most 5 * 10^5.
*/ 

package main

type TrieNode struct {
	children  [26]*TrieNode
	bestIndex int
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &TrieNode{bestIndex: 0}

	isBetter := func(idxNew, idxOld int) bool {
		lenNew := len(wordsContainer[idxNew])
		lenOld := len(wordsContainer[idxOld])
		if lenNew < lenOld {
			return true
		}
		if lenNew == lenOld && idxNew < idxOld {
			return true
		}
		return false
	}

	// 1. Start from Root
	for i := 1; i < len(wordsContainer); i++ {
		if isBetter(i, root.bestIndex) {
			root.bestIndex = i
		}
	}

	// 2. Input word from wordsContainer in Trie
	for i, word := range wordsContainer {
		curr := root
		for j := len(word) - 1; j >= 0; j-- {
			charIdx := word[j] - 'a'
			if curr.children[charIdx] == nil {
				curr.children[charIdx] = &TrieNode{bestIndex: i}
			}
			curr = curr.children[charIdx]

			if isBetter(i, curr.bestIndex) {
				curr.bestIndex = i
			}
		}
	}

	// 3. Find ans from query
	ans := make([]int, len(wordsQuery))
	for i, query := range wordsQuery {
		curr := root
		for j := len(query) - 1; j >= 0; j-- {
			charIdx := query[j] - 'a'
			if curr.children[charIdx] == nil {
				break
			}
			curr = curr.children[charIdx]
		}
		ans[i] = curr.bestIndex
	}

	return ans
}


