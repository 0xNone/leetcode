package leetcode

// fix: BastTest3，end = 0 跳过 0 的情况
func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)
	res := []int{-1, -1}
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			if res[0] == -1 || mid < res[0] {
				res[0] = mid
			}
			end = mid
		}
	}
	start, end = 0, len(nums)
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			if res[1] == -1 || mid > res[1] {
				res[1] = mid
			}
			start = mid + 1
		}
	}
	return res
}
