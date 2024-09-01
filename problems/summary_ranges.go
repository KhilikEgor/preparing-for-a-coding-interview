// https://leetcode.com/problems/summary-ranges/
package problems

import "fmt"


func summaryRanges(nums []int) []string {
    var ans []string

    for i := 0; i < len(nums); i++ {
        start := nums[i]
        for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
            i++
        }
        if start == nums[i] {
            ans = append(ans, fmt.Sprintf("%d", start))
        } else {
            ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i]))
        }
    }
    return ans
}