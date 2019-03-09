package leetcode

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"BaseTest", args{"2", "3"}, "6"},
		{"BaseTest", args{"123", "456"}, "56088"},
		{"StuckTest", args{"0", "9"}, "0"},
		{"StuckTest", args{"498828660196", "840477629533"}, "419254329864656431168468"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
