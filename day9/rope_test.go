package day9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func Test_rope_moveTail(t *testing.T) {
	l := zerolog.Nop()
	ttrace := make(map[int]struct{})

	tests := [][2]int{
		{-2, -2},
		{-1, -2},
		{-0, -2},
		{1, -2},
		{2, -2},

		{-2, -1},
		{-1, -1},
		{-0, -1},
		{1, -1},
		{2, -1},

		{-2, 0},
		{-1, 0},
		{-0, 0},
		{1, 0},
		{2, 0},

		{-2, 1},
		{-1, 1},
		{-0, 1},
		{1, 1},
		{2, 1},

		{-2, 2},
		{-1, 2},
		{-0, 2},
		{1, 2},
		{2, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			r := &rope{
				l:         l,
				head:      [2]int{0, 0},
				tail:      tt,
				tailTrace: ttrace,
			}

			t.Logf("before: \n%s", visualizeHT(r.head, r.tail))

			if err := r.moveTail(); (err != nil) != false {
				t.Errorf("moveTail() error = %v, wantErr %v", err, false)
			}

			t.Logf("after: \n\n%s\n", visualizeHT(r.head, r.tail))
		})
	}
}

func visualizeHT(h, t [2]int) string {
	upDown := h[0] - t[0]
	leftRight := h[1] - t[1]
	var sb strings.Builder
	char := "."
	for row := -2; row < 3; row++ {
		for col := -2; col < 3; col++ {
			switch {
			case row == 0 && col == 0:
				char = "H"
			case row == upDown && col == leftRight:
				char = "T"
			default:
				char = "."
			}

			sb.WriteString(char)
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func Test_coordToString(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0,0",
			args: args{
				x: 0,
				y: 0,
			},
			want: "0 - 0",
		},
		{
			name: "300,221",
			args: args{
				x: 300,
				y: 221,
			},
			want: "300 - 221",
		},
		{
			name: "-93,13",
			args: args{
				x: -93,
				y: 13,
			},
			want: "-93 - 13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coordToString(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("coordToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
