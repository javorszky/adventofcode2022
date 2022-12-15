package day15

import (
	"testing"
)

func Test_line_isCoordInLine(t *testing.T) {
	type fields struct {
		start coordinate
		end   coordinate
	}
	type args struct {
		c coordinate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "horizontal, is in line",
			fields: fields{
				start: coordinate{0, 6},
				end:   coordinate{10, 6},
			},
			args: args{c: coordinate{5, 6}},
			want: true,
		},
		{
			name: "horizontal, is not in line (different plane)",
			fields: fields{
				start: coordinate{0, 6},
				end:   coordinate{10, 6},
			},
			args: args{c: coordinate{5, 5}},
			want: false,
		},
		{
			name: "horizontal, is not in line (out of bounds)",
			fields: fields{
				start: coordinate{0, 6},
				end:   coordinate{10, 6},
			},
			args: args{c: coordinate{12, 6}},
			want: false,
		},

		{
			name: "vertical, is in line",
			fields: fields{
				start: coordinate{6, 0},
				end:   coordinate{6, 10},
			},
			args: args{c: coordinate{6, 5}},
			want: true,
		},
		{
			name: "vertical, is not in line (different plane)",
			fields: fields{
				start: coordinate{6, 0},
				end:   coordinate{6, 10},
			},
			args: args{c: coordinate{5, 5}},
			want: false,
		},
		{
			name: "horizontal, is not in line (out of bounds)",
			fields: fields{
				start: coordinate{6, 0},
				end:   coordinate{6, 10},
			},
			args: args{c: coordinate{6, 12}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := newLine(tt.fields.start, tt.fields.end)
			if err != nil {
				t.Errorf("could not create line for %v, %v: %s", tt.fields.start, tt.fields.end, err.Error())
			}

			if got := l.isCoordInLine(tt.args.c); got != tt.want {
				t.Errorf("isCoordInLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
