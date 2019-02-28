package leetcode

// fix: BaseTest3，单个 0 的情况
// fix: BaseTest4
// fix: BaseTest5，单数 1 计算错误
func isOneBitCharacter(bits []int) bool {
	bitsLen := len(bits)
	if bitsLen == 1 && bits[bitsLen-1] == 0 {
		return true
	}
	if bits[bitsLen-1] != 0 {
		return false
	}
	oneNums := 0
	if bits[bitsLen-2] == 1 {
		for i := bitsLen - 2; i >= 0; i-- {
			if bits[i] == 1 {
				oneNums += 1
			} else {
				break
			}
		}
	}
	if oneNums%2 != 0 {
		return false
	}
	return true
}
