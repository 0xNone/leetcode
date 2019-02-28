package leetcode

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"BaseTest", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{[]string{"ate", "eat", "tea"}, []string{"nat", "tan"}, []string{"bat"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); len(got) != len(tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
