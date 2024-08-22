// https://leetcode.com/problems/move-zeroes/description/?envType=study-plan-v2&envId=leetcode-75
package problems

func moveZeroes(nums []int)  {
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0{
			nums[ptr] = nums[i]
			ptr++
		}
	}

	for ptr < len(nums){
		nums[ptr] = 0
		ptr++
	}
}