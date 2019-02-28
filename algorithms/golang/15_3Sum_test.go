package leetcode

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"BaseTest", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
		{"BaseTest", args{[]int{1, -1, -1, 0}}, [][]int{[]int{-1, 0, 1}}},
		{"BaseTest", args{[]int{0, 0, 0, 0}}, [][]int{[]int{0, 0, 0}}},
		{"BaseTest", args{[]int{1, 1, -2}}, [][]int{[]int{-2, 1, 1}}},
		{"BaseTest", args{[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}}, [][]int{[]int{-5, 1, 4}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-2, -2, 4}, []int{-2, 1, 1}, []int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); len(got) != len(tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
