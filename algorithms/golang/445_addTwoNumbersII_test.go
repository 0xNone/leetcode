package leetcode

import (
	"testing"
)

func Test_addTwoNumbersii(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"BaseTest", args{&ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}}, &ListNode{7, &ListNode{8, &ListNode{0, &ListNode{7, nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbersii(tt.args.l1, tt.args.l2)
			for true {
				if tt.want == nil && got == nil {
					break
				} else if tt.want == nil || got == nil || tt.want.Val != got.Val {
					t.Errorf("addTwoNumbersii() = %v, want %v", got, tt.want)
				}
				tt.want = tt.want.Next
				got = got.Next
			}
		})
	}
}
