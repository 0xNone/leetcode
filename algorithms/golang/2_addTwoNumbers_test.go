package leetcode

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"BaseTest", args{&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
		{"StuckTest", args{&ListNode{5, nil}, &ListNode{5, nil}}, &ListNode{0, &ListNode{1, nil}}},
		{"StuckTest", args{&ListNode{1, &ListNode{8, nil}}, &ListNode{0, nil}}, &ListNode{1, &ListNode{8, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			for true {
				if tt.want == nil && got == nil {
					break
				} else if tt.want == nil || got == nil || tt.want.Val != got.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
				}
				tt.want = tt.want.Next
				got = got.Next
			}
		})
	}
}
