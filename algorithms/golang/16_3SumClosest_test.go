package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"BaseTest", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"BaseTest", args{[]int{1, 1, 1, 1}, 0}, 3},
		{"BaseTest", args{[]int{1, 1, 1, 0}, -100}, 2},
		{"BaseTest", args{[]int{1, 2, 4, 8, 16, 32, 64, 128}, 82}, 82},
		{"BaseTest", args{[]int{1, 2, 5, 10, 11}, 12}, 13},
		{"BaseTest", args{[]int{0, 2, 1, -3}, 1}, 0},
		{"BaseTest", args{[]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
