package leetcode

//https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=programming-skills

import "fmt"

func MergeAlternately(word1 string, word2 string) string {
	var merge string
	n := max(len(word1), len(word2))

	for i := 0; i < n; i++ {
		if i < len(word1) {
			merge = merge + string(word1[i])
			fmt.Println(merge, len(merge)) //ver word1 em partes
		}
		if i < len(word2) {
			merge = merge + string(word2[i])
			fmt.Println(merge, len(merge)) //ver word2 em partes
		}

	}
	return merge
}

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
