// https://leetcode.com/problems/squares-of-a-sorted-array/description/
package problems

func sortedSquares(nums []int) []int {
    ans := make([]int, len(nums))
    len_array := len(nums) - 1
    prt1, prt2 := 0, len(nums) - 1

    for prt1 <= prt2 {
        ptr1_square := nums[prt1] * nums[prt1]
        ptr2_square := nums[prt2] * nums[prt2]
        if ptr1_square < ptr2_square {
            ans[len_array] = ptr2_square
            prt2--
            len_array--
        } else {
            ans[len_array] = ptr1_square
            prt1++
            len_array--
        }
    }
    return ans
}

// [-4,-1,0,3,10]

// first way
// [16, 1, 0, 9, 100]  
// [0, 1, 9, 16, 100]  n^2

//second way
// ans: [0 ,1 ,9 , 16, 100]
// [-4,-1,0,3,10]
//      ^   ^
//      1   9  
// O(n) - память
// O(n) - скорость
