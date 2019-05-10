package leetcode

import "sort"

func frequencySort(s string) string {
	charMap := map[rune]int{}
	type keyValue struct {
		Key   rune
		Value int
	}
	res := []rune{}
	for _, v := range s {
		if _, ok := charMap[v]; ok {
			charMap[v]++
		} else {
			charMap[v] = 1
		}
	}

	kvSlice := []keyValue{}
	for k, v := range charMap {
		kvSlice = append(kvSlice, keyValue{k, v})
	}
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	for i := range kvSlice {
		for j := 0; j < kvSlice[i].Value; j++ {
			res = append(res, kvSlice[i].Key)
		}
	}
	return string(res)
}
