package leetcode

// todo: 大佬写的，太牛逼了，观察推算了30分钟
func find132pattern(nums []int) bool {
	stack := []int{}
	first := true
	min := 0
	for i := len(nums) - 1; i > -1; i-- {
		if nums[i] < min && !first {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			first = false
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
