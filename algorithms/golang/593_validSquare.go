package leetcode

func FormIsRightAngle(p1, p2, p3 []int) bool {
	a := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	b := (p2[0]-p3[0])*(p2[0]-p3[0]) + (p2[1]-p3[1])*(p2[1]-p3[1])
	c := (p3[0]-p1[0])*(p3[0]-p1[0]) + (p3[1]-p1[1])*(p3[1]-p1[1])
	if a > b && a > c {
		if a == b+c && b == c {
			return true
		}
	} else if b > a && b > c {
		if b == a+c && a == c {
			return true
		}
	} else {
		if c == a+b && a == b {
			return true
		}
	}
	return false
}

// fix: BaseTest2，边长问题
// fix: BaseTest3，重复点问题
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if IsSameIntSlice(p1, p2) || IsSameIntSlice(p2, p3) || IsSameIntSlice(p3, p4) || IsSameIntSlice(p4, p1) {
		return false
	}
	if FormIsRightAngle(p1, p2, p3) && FormIsRightAngle(p2, p3, p4) && FormIsRightAngle(p3, p4, p1) && FormIsRightAngle(p4, p1, p2) {
		return true
	}
	return false
}
