// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element
package problem

func longestSubarray(nums []int) int {
	left, right, zeroCount, maxLen := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
		right++
	}
	return maxLen
}