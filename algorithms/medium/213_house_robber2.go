package medium

import "github.com/leetcodego/utils"

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}
	return utils.Max(subRob(nums[:len(nums)-2]), subRob(nums[2:]))
}

func subRob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	temp1, temp2 := nums[0], utils.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		temp1, temp2 = temp2, utils.Max(temp1+nums[i], temp2)
	}
	return temp2
}
