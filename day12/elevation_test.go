package day12

import (
	"testing"
)

func Test_coordToBinary(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0, 0",
			args: args{
				row: 0,
				col: 0,
			},
			want: 0,
		},
		{
			name: "10, 10",
			args: args{
				row: 10,
				col: 10,
			},
			want: 0b0000101000001010,
		},
		{
			name: "64, 128",
			args: args{
				row: 64,
				col: 128,
			},
			want: 0b0100000010000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coordToBinary(tt.args.row, tt.args.col)
			if got != tt.want {
				t.Errorf("coordToBinary() = %016b, want %016b", got, tt.want)
			}

			gotRow, gotCol := binaryToCoord(got)
			if gotRow != tt.args.row {
				t.Errorf("binaryToCoord() row = %d, want %d", gotRow, tt.args.row)
			}
			if gotCol != tt.args.col {
				t.Errorf("binaryToCoord() col = %d, want %d", gotCol, tt.args.col)
			}
		})
	}
}

func Test_moveUp(t *testing.T) {
	type want struct {
		up, left, down, right [2]int
	}

	tests := []struct {
		name   string
		rowCol [2]int
		want   want
	}{
		{
			name:   "can move correctly from point",
			rowCol: [2]int{20, 20},
			want: want{
				up:    [2]int{19, 20},
				left:  [2]int{20, 19},
				down:  [2]int{21, 20},
				right: [2]int{20, 21},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coordToBinary(tt.rowCol[0], tt.rowCol[1])

			u, u2 := binaryToCoord(moveUp(c))
			l, l2 := binaryToCoord(moveLeft(c))
			d, d2 := binaryToCoord(moveDown(c))
			r, r2 := binaryToCoord(moveRight(c))

			if [2]int{u, u2} != tt.want.up {
				t.Errorf("moveUp()= %v, want %v", [2]int{u, u2}, tt.want.up)
			}

			if [2]int{l, l2} != tt.want.left {
				t.Errorf("moveLeft() = %v, want %v", [2]int{l, l2}, tt.want.left)
			}

			if [2]int{d, d2} != tt.want.down {
				t.Errorf("moveDown() = %v, want %v", [2]int{d, d2}, tt.want.down)
			}

			if [2]int{r, r2} != tt.want.right {
				t.Errorf("moveRight() = %v, want %v", [2]int{r, r2}, tt.want.up)
			}
		})
	}
}
