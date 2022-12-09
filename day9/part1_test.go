package day9

import "testing"

func Test_coordToBinary(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "turns 255, 255 into binary",
			args: args{
				x: 255,
				y: 255,
			},
			want: 0b10111111111011111111,
		},
		{
			name: "turns -255, -255 into binary",
			args: args{
				x: -255,
				y: -255,
			},
			want: 0b01000000010100000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coordToBinary(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("coordToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1 2 4 8
// 16 32 64 128
//
