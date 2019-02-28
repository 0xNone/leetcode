package leetcode

import (
	"fmt"
	"strings"
)

func StringHandler(s string) []string {
	res := []string{}
	sLen := len(s)
	if sLen < 2 {
		res = append(res, s)
		return res
	}
	j := 0
	for ; j < sLen; j++ {
		if s[j] != '0' {
			break
		}
	}
	if j == sLen {
		res = append(res, "0")
	} else {
		endIndex := sLen - 1
		for ; endIndex > 0; endIndex-- {
			if s[endIndex] != '0' {
				break
			}
		}
		res = append(res, s[j:])
		if j != 0 {
			res = append(res, "0."+s[j:endIndex+1])
		}
		for k := j + 1; k <= endIndex; k++ {
			res = append(res, s[j:k]+"."+s[k:endIndex+1])
		}
	}
	return res
}

// fix: BaseTest2，题目看错了，要求所有数字都存在
func ambiguousCoordinates(S string) []string {
	numS := S[1 : len(S)-1]
	res := []string{}
	sLen := len(numS)
	for i := 0; i <= sLen; i++ {
		if i < 1 || i > sLen-1 {
			continue
		}
		l := numS[:i]
		r := numS[i:]
		ls := []string{}
		rs := []string{}
		if len(l) == 1 {
			ls = append(ls, l)
		} else {
			allZero := true
			for li := range l {
				if l[li] != '0' {
					allZero = false
					break
				}
			}
			if allZero {
				continue
			}
			for li := range l {
				if li == 0 {
					if l[li] == '0' {
						if strings.HasSuffix(l, "0") {
							break
						}
						ls = append(ls, "0."+l[1:])
						break
					} else {
						ls = append(ls, l)
					}
					continue
				}
				if l[len(l)-1] == '0' {
					break
				}
				ls = append(ls, l[:li]+"."+l[li:])
			}
		}
		if len(r) == 1 {
			rs = append(rs, r)
		} else {
			allZero := true
			for ri := range r {
				if r[ri] != '0' {
					allZero = false
					break
				}
			}
			if allZero {
				continue
			}
			for ri := range r {
				if ri == 0 {
					if r[ri] == '0' {
						if strings.HasSuffix(r, "0") {
							break
						}
						rs = append(rs, "0."+r[1:])
						break
					} else {
						rs = append(rs, r)
					}
					continue
				}
				if strings.HasSuffix(r, "0") {
					break
				}
				rs = append(rs, r[:ri]+"."+r[ri:])
			}
		}
		for _, v1 := range ls {
			for _, v2 := range rs {
				res = append(res, fmt.Sprintf("(%v, %v)", v1, v2))
			}
		}
	}
	return res
}
