package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
			if i == 0 {
				res := []int{}
				res = append(res, 1)
				for i := range digits {
					res = append(res, digits[i])
				}
				digits = res
			}
		}
	}
	return digits
}
