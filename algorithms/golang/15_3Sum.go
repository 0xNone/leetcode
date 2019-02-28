package leetcode

import (
	"sort"
)

// 修复 baseTest4，不应该判断 nums[j] > 0 && nums[k] > 0 就 break
//func threeSum(nums []int) [][]int {
//	res := [][]int{}
//	sort.Ints(nums)
//	for i, v := range nums {
//		if v > 0 {
//			break
//		}
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		j := i + 1
//		k := len(nums) - 1
//		for j < len(nums)-1 {
//			for j < k {
//				//if nums[j] > 0 && nums[k] > 0 {
//				//	break
//				//}
//				if nums[j]+nums[k] == -v {
//					res = append(res, []int{v, nums[j], nums[k]})
//				}
//				for j < k && nums[k] == nums [k-1] {
//					k--
//				}
//				k--
//			}
//			k = len(nums) - 1
//			for j < k && nums[j] == nums[j+1] {
//				j++
//			}
//			j++
//		}
//	}
//	return res
//}

// 看网友思路编写，二分法，理论时间复杂度应为 O(log n)
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := v + nums[l] + nums[r]
			if s == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
