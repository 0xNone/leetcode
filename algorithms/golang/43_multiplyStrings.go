package leetcode

// fix: StuckTest1，未考虑结果为 0 的情况
// fix: StuckTest2，重写，结果很大的情况
func multiply(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	res := ""

	n1Len := len(num1)
	n2Len := len(num2)
	tRes := make([]uint8, n1Len+n2Len)
	for i := n1Len - 1; i >= 0; i-- {
		for j := n2Len - 1; j >= 0; j-- {
			t := (num1[i] - '0') * (num2[j] - '0')
			tRes[i+j+1] += t
			if tRes[i+j+1]/10 > 0 {
				tRes[i+j] += tRes[i+j+1] / 10
				tRes[i+j+1] %= 10
			}
		}
	}

	startNum := 0
	start := false
	for i := 0; i < n1Len+n2Len; i++ {
		if tRes[i] != 0 && !start {
			startNum = i
			start = true
		}
		res += string(tRes[i] + '0')
	}
	if !start {
		return "0"
	}
	return res[startNum:]
}

//func multiply(num1 string, num2 string) string {
//	if num1 == "0" && num2 == "0" {
//		return "0"
//	}
//	res := ""
//	var iNum int = 0
//	numLen := len(num1)
//	for i := 0; i < numLen; i++ {
//		iNum += int(num1[i]-'0') * int(math.Pow10(numLen-i-1))
//	}
//
//	iRes := 0
//	numLen = len(num2)
//	for i := numLen - 1; i >= 0; i-- {
//		iRes += int(num2[i]-'0') * iNum * int(math.Pow10(numLen-i-1))
//	}
//	if iRes == 0 {
//		return "0"
//	}
//	for iRes != 0 {
//		res = string(iRes%10+'0') + res
//		iRes /= 10
//	}
//	return res
//}
