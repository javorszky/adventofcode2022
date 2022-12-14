package day14

import (
	"reflect"
	"testing"
)

func Test_xyToBinary(t *testing.T) {
	tests := []struct {
		name string
		x, y int
	}{
		{
			name: "0,0",
			x:    0,
			y:    0,
		},
		{
			name: "2047, 2047",
			x:    2047,
			y:    2047,
		},
		{
			name: "918, 23",
			x:    918,
			y:    23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bin := xyToBinary(tt.x, tt.y)
			gotX, gotY := binaryToXY(bin)
			if tt.x != gotX {
				t.Errorf("x to bin to x is borked: got %d, want %d", gotX, tt.x)
			}

			if tt.y != gotY {
				t.Errorf("y to bin to y is borked: got %d, want %d", gotY, tt.y)
			}
		})
	}
}

func Test_generateConnectors(t *testing.T) {
	type args struct {
		a [2]int
		b [2]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][2]int
		wantErr bool
	}{
		{
			name: "x is same, y1 is smaller",
			args: args{
				a: [2]int{3, 7},
				b: [2]int{3, 9},
			},
			want: [][2]int{
				{3, 7},
				{3, 8},
				{3, 9},
			},
		},
		{
			name: "x is same, y2 is smaller",
			args: args{
				a: [2]int{3, 7},
				b: [2]int{3, 4},
			},
			want: [][2]int{
				{3, 4},
				{3, 5},
				{3, 6},
				{3, 7},
			},
		},
		{
			name: "y is same, x1 is smaller",
			args: args{
				a: [2]int{3, 4},
				b: [2]int{6, 4},
			},
			want: [][2]int{
				{3, 4},
				{4, 4},
				{5, 4},
				{6, 4},
			},
		},
		{
			name: "y is same, x2 is smaller",
			args: args{
				a: [2]int{3, 4},
				b: [2]int{0, 4},
			},
			want: [][2]int{
				{0, 4},
				{1, 4},
				{2, 4},
				{3, 4},
			},
		},
		{
			name: "neither are the same",
			args: args{
				a: [2]int{3, 4},
				b: [2]int{0, 5},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateConnectors(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateConnectors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateConnectors() got = %v, want %v", got, tt.want)
			}
		})
	}
}
