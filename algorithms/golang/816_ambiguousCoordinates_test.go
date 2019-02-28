package leetcode

import (
	"testing"
)

func Test_ambiguousCoordinates(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"BaseTest", args{"(123)"}, []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"}},
		{"BaseTest", args{"(00011)"}, []string{"(0.001, 1)", "(0, 0.011)"}},
		{"BaseTest", args{"(0123)"}, []string{"(0, 1.23)", "(0, 12.3)", "(0, 123)", "(0.1, 2.3)", "(0.1, 23)", "(0.12, 3)"}},
		{"BaseTest", args{"(100)"}, []string{"(10, 0)"}},
		{"BaseTest", args{"(0010)"}, []string{"(0.01, 0)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ambiguousCoordinates(tt.args.S); len(got) != len(tt.want) {
				t.Errorf("ambiguousCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
