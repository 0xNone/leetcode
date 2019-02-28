package leetcode

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"BaseTest", args{[]int{1, 0, 0}}, true},
		{"BaseTest", args{[]int{1, 1, 1, 0}}, false},
		{"BaseTest", args{[]int{0}}, true},
		{"BaseTest", args{[]int{1, 1, 0}}, true},
		{"BaseTest", args{[]int{1, 0, 1, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
