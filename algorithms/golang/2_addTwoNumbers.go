package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// fix: StuckTest1，进位问题
// fix: StuckTest2，l1 和 l2 长度不等
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	resPtr := res
	for true {
		if l1 == nil {
			resPtr.Val += l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			resPtr.Val += l1.Val
			l1 = l1.Next
		} else {
			resPtr.Val += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		tNode := &ListNode{resPtr.Val / 10, nil}
		if tNode.Val != 0 {
			resPtr.Val %= 10
		}
		if l1 == nil && l2 == nil {
			if tNode.Val != 0 {
				resPtr.Next = tNode
			}
			break
		}
		resPtr.Next = tNode
		resPtr = resPtr.Next
	}
	return res
}
