package leetcode

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	sort.Strings(strs)
	resMap := map[string][]string{}
	for i := range strs {
		sortStr := SortString(strs[i])
		if _, ok := resMap[sortStr]; ok {
			resMap[sortStr] = append(resMap[sortStr], strs[i])
		} else {
			resMap[sortStr] = []string{strs[i]}
		}
	}

	res := [][]string{}
	for _, v := range resMap {
		res = append(res, v)
	}
	return res
}
