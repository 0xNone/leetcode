package leetcode

import (
	"sort"
)

func IntAbs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

// fix: BaseTest2，未考虑 nums 全部大于 0 的情况
// fix: BaseTest3，尽管 res-target >= 0 时 i 还往右走的情况
// fix: BaseTest4，计算内容的情况
// fix: BaseTest5，重新调整最小差值计算
func threeSumClosest(nums []int, target int) int {
	closeDiff := 0
	sort.Ints(nums)
	closeDiff = nums[0] + nums[1] + nums[2] - target
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			res := nums[i] + nums[l] + nums[r]
			if IntAbs(res-target) < IntAbs(closeDiff) {
				closeDiff = res - target
			}
			if res-target < 0 {
				l++
			} else if res-target > 0 {
				r--
			} else {
				break
			}
		}
	}
	return closeDiff + target
}
