package leetcode

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"baseTest", args{"tree"}, "eetr"},
		{"baseTest", args{"cccaaa"}, "cccaaa"},
		{"baseTest", args{"Aabb"}, "bbAa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
