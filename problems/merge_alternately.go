// https://leetcode.com/problems/merge-strings-alternately/

package problems

import "strings"

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func mergeAlternately(word1 string, word2 string) string {
    var merged strings.Builder
    len1, len2 := len(word1), len(word2)
    minLen :=  min(len1, len2)

    for i := 0; i < minLen; i++ {
        merged.WriteByte(word1[i])
        merged.WriteByte(word2[i])
    }

    merged.WriteString(word1[minLen:])
    merged.WriteString(word2[minLen:])

    return merged.String()
}