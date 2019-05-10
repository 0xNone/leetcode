package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	type keyValue struct {
		Key   int
		Value int
	}
	res := []int{}
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v]++
		} else {
			numMap[v] = 1
		}
	}

	kvSlice := []keyValue{}
	for k, v := range numMap {
		kvSlice = append(kvSlice, keyValue{k, v})
	}
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	for i := 0; i < k; i++ {
		res = append(res, kvSlice[i].Key)
	}
	return res
}
