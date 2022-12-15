package day15

import (
	"os"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
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

func Test_cutLines(t *testing.T) {
	l := zerolog.New(os.Stderr).With().Str("module", "cutlinestest").Logger()

	type args struct {
		origin line
		cutSet lines
	}
	tests := []struct {
		name string
		args args
		want lines
	}{
		{
			name: "cutset out of plane",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{5, 11},
						end:         coordinate{7, 11},
						orientation: lineHorizontal,
						rowCol:      11,
					},
				},
			},
			want: lines{
				{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
			},
		},
		{
			name: "cutset from tail end",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{5, 10},
						end:         coordinate{25, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{
				line{
					start:       coordinate{0, 10},
					end:         coordinate{4, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
			},
		},
		{
			name: "cutset from leading end",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{-5, 10},
						end:         coordinate{5, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{
				{
					start:       coordinate{6, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
			},
		},
		{
			name: "cutset envelops",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{-1, 10},
						end:         coordinate{21, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{},
		},
		{
			name: "cutset matches",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{0, 10},
						end:         coordinate{20, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{},
		},
		{
			name: "cutset contained",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{5, 10},
						end:         coordinate{10, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{
				{
					start:       coordinate{0, 10},
					end:         coordinate{4, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				{
					start:       coordinate{11, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
			},
		},
		{
			name: "multiple cutsets",
			args: args{
				origin: line{
					start:       coordinate{0, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				cutSet: lines{
					{
						start:       coordinate{2, 10},
						end:         coordinate{4, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
					{
						start:       coordinate{9, 10},
						end:         coordinate{11, 10},
						orientation: lineHorizontal,
						rowCol:      10,
					},
				},
			},
			want: lines{
				{
					start:       coordinate{0, 10},
					end:         coordinate{1, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				{
					start:       coordinate{5, 10},
					end:         coordinate{8, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
				{
					start:       coordinate{12, 10},
					end:         coordinate{20, 10},
					orientation: lineHorizontal,
					rowCol:      10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cutLines(tt.args.origin, tt.args.cutSet, l)
			if err != nil {
				t.Errorf("got an error, and that shouldn't have happened: %s", err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
