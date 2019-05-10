package leetcode

// fix: errorTest1，没有考虑 ' 的情况
// fix: errorTest2，没有考虑 - 的情况
// fix: errorTest3，em...只需要考虑非空格既是单词即可
func countSegments(s string) int {
	detected := true
	wordCount := 0
	for i := range s {
		if s[i] != ' ' {
			if detected {
				wordCount += 1
			}
			detected = false
		} else {
			detected = true
		}
		//if (s[i] <= 122 && s[i] >= 97) || (s[i] <= 90 && s[i] >= 65) || s[i] == 39 || s[i] == 45 {
		//	if detected {
		//		wordCount += 1
		//	}
		//	detected = false
		//} else {
		//	detected = true
		//}
	}
	return wordCount
}
