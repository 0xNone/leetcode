package leetcode

func SliceIntElementsEqual(v1, v2 []int) bool {
	v1m := map[int]int{}
	for _, v := range v1 {
		if _, ok := v1m[v]; !ok {
			v1m[v] = 1
		} else {
			v1m[v] += 1
		}
	}
	v2m := map[int]int{}
	for _, v := range v2 {
		if _, ok := v2m[v]; !ok {
			v2m[v] = 1
		} else {
			v2m[v] += 1
		}
	}

	for k, v1 := range v1m {
		if v2, ok := v2m[k]; !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}

func IsSameIntSlice(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

func IntContains(s int, t []int) bool {
	for _, v := range t {
		if v == s {
			return true
		}
	}
	return false
}
