package leetcode

func addTwoNumbersii(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	s1 := []int{}
	s2 := []int{}
	for true {
		if l1 == nil {
			break
		}
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for true {
		if l2 == nil {
			break
		}
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	s1Len := len(s1)
	s2Len := len(s2)
	for i := 1; true; {
		if i > s1Len {
			res.Val += s2[s2Len-i]
		} else if i > s2Len {
			res.Val += s1[s1Len-i]
		} else {
			res.Val += s1[s1Len-i] + s2[s2Len-i]
		}
		tNode := &ListNode{res.Val / 10, res}
		if res.Val > 9 {
			res.Val %= 10
		}
		i++
		if i > s1Len && i > s2Len {
			if tNode.Val != 0 {
				res = tNode
			}
			break
		}
		res = tNode
	}
	return res
}
