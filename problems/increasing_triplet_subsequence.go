// https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75
package problems


func increasingTriplet(nums []int) bool {

    first, second := math.MaxInt32, math.MaxInt32

    for _, num := range nums {
        if num <= first {
            first = num
        } else if num <= second {
            second = num
        } else {
            return true
        }
    }

    return false
}

//[2,1,5,0,4,6]
//       
// f := 99 | 2 1 0
// s := 99 | 5 4
// Гениально