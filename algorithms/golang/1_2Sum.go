package leetcode

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if v1 + v2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//func twoSum(nums []int, target int) []int {
//	t := map[int]int{}
//	for i, v := range nums {
//		res := target - v
//		if index, ok := t[res]; ok && i != index {
//			return []int{index, i}
//		}
//		t[v] = i
//	}
//	return nil
//}