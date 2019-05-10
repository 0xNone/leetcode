package leetcode

import "testing"

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"baseTest", args{"Hello, my name is John"}, 5},
		{"errorTest", args{"love live! mu'sic forever"}, 4},
		{"errorTest", args{"The one-hour drama series Westworld is a dark odyssey about the dawn of artificial consciousness and the evolution of sin. Set at the intersection of the near future and the reimagined past, it explores a world in which every human appetite, no matter how noble or depraved, can be indulged."}, 50},
		{"errorTest", args{", , , ,        a, eaefa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
