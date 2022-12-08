package day8

import (
	"testing"
)

func Test_coordToByte(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "converts 0,0 to 0",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
		{
			name: "converts 8, 8 to something",
			args: args{
				x: 8,
				y: 8,
			},
			want: 0b0000100000001000,
		},
		{
			name: "converts 127,127 to something",
			args: args{
				x: 127,
				y: 127,
			},
			want: 0b0111111101111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coordToBinary(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("coordToByte() = %v, %016b, want %v", got, got, tt.want)
			}
		})
	}
}
