package day15

import (
	"reflect"
	"testing"
)

func Test_sensor_rowInExclusion(t *testing.T) {
	s := newSensor(coordinate{10, 10}, coordinate{12, 12})

	tests := []struct {
		name string
		row  int
		want bool
	}{
		{
			name: "row goes through sensor, in exclusion",
			row:  10,
			want: true,
		},
		{
			name: "row is at top tip, in exclusion",
			row:  6,
			want: true,
		},
		{
			name: "row is just above top tip, not in exclusion",
			row:  5,
			want: false,
		},
		{
			name: "row is bottom tip, in exclusion",
			row:  14,
			want: true,
		},
		{
			name: "row is just below bottom tip, not in exclusion",
			row:  15,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.rowInExclusion(tt.row); got != tt.want {
				t.Errorf("rowInExclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensor_rowBoundLine(t *testing.T) {
	s := newSensor(coordinate{10, 10}, coordinate{12, 12})

	tests := []struct {
		name    string
		row     int
		want    line
		wantErr bool
	}{
		{
			name: "row goes through sensor",
			row:  10,
			want: line{
				start:       coordinate{6, 10},
				end:         coordinate{14, 10},
				orientation: lineHorizontal,
				rowCol:      10,
			},
			wantErr: false,
		},
		{
			name: "row is top tip",
			row:  6,
			want: line{
				start:       coordinate{10, 6},
				end:         coordinate{10, 6},
				orientation: lineHorizontal,
				rowCol:      6,
			},
			wantErr: false,
		},
		{
			name: "row is somewhere in the middle",
			row:  8,
			want: line{
				start:       coordinate{8, 8},
				end:         coordinate{12, 8},
				orientation: lineHorizontal,
				rowCol:      8,
			},
			wantErr: false,
		},
		{
			name:    "row does not intersect",
			row:     0,
			want:    line{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.lineForRow(tt.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("lineForRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lineForRow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newLine(t *testing.T) {
	type args struct {
		a coordinate
		b coordinate
	}
	tests := []struct {
		name    string
		args    args
		want    line
		wantErr bool
	}{
		{
			name: "does not make a line, would be at an angle",
			args: args{
				a: coordinate{2, 3},
				b: coordinate{9, 23},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "makes a horizontal line, coords right order",
			args: args{
				a: coordinate{0, 5},
				b: coordinate{10, 5},
			},
			want: line{
				start: coordinate{0, 5},
				end:   coordinate{10, 5},
			},
			wantErr: false,
		},
		{
			name: "makes a horizontal line, coords flipped order",
			args: args{
				a: coordinate{10, 5},
				b: coordinate{0, 5},
			},
			want: line{
				start: coordinate{0, 5},
				end:   coordinate{10, 5},
			},
			wantErr: false,
		},
		{
			name: "makes a vertical line, coords right order",
			args: args{
				a: coordinate{5, 0},
				b: coordinate{5, 10},
			},
			want: line{
				start: coordinate{5, 0},
				end:   coordinate{5, 10},
			},
			wantErr: false,
		},
		{
			name: "makes a vertical line, coords flipped order",
			args: args{
				a: coordinate{5, 10},
				b: coordinate{5, 0},
			},
			want: line{
				start: coordinate{5, 0},
				end:   coordinate{5, 10},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newLine(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("newLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeLines(t *testing.T) {
	type args struct {
		a line
		b line
	}
	tests := []struct {
		name    string
		args    args
		want    line
		wantErr bool
	}{
		{
			name: "lines different orientation",
			args: args{
				a: line{
					start:       coordinate{5, 0},
					end:         coordinate{5, 10},
					orientation: lineVertical,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{0, 5},
					end:         coordinate{10, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines not on same plane, horizontal",
			args: args{
				a: line{
					start:       coordinate{0, 5},
					end:         coordinate{10, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{2, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines not on same plane, vertical",
			args: args{
				a: line{
					start:       coordinate{5, 0},
					end:         coordinate{5, 10},
					orientation: lineVertical,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{6, 2},
					end:         coordinate{6, 12},
					orientation: lineVertical,
					rowCol:      6,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines do not touch horizontal, good order",
			args: args{
				a: line{
					start:       coordinate{0, 5},
					end:         coordinate{10, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{12, 5},
					end:         coordinate{17, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines do not touch horizontal, flipped order",
			args: args{
				a: line{
					start:       coordinate{12, 5},
					end:         coordinate{17, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{0, 5},
					end:         coordinate{10, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines do not touch vertical, good order",
			args: args{
				a: line{
					start:       coordinate{5, 0},
					end:         coordinate{5, 10},
					orientation: lineVertical,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{5, 12},
					end:         coordinate{5, 17},
					orientation: lineVertical,
					rowCol:      5,
				},
			},
			want:    line{},
			wantErr: true,
		},
		{
			name: "lines do not touch vertical, flipped order",
			args: args{
				a: line{
					start:       coordinate{5, 12},
					end:         coordinate{5, 17},
					orientation: lineVertical,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{5, 0},
					end:         coordinate{5, 10},
					orientation: lineVertical,
					rowCol:      5,
				},
			},
			want:    line{},
			wantErr: true,
		},

		{
			name: "lines touch horizontal, good order",
			args: args{
				a: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{12, 5},
					end:         coordinate{17, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{17, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},
		{
			name: "lines touch horizontal, flipped order",
			args: args{
				a: line{
					start:       coordinate{12, 5},
					end:         coordinate{17, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{17, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},
		{
			name: "lines engulf horizontal, good order",
			args: args{
				a: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{2, 5},
					end:         coordinate{7, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{11, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},
		{
			name: "lines engulf horizontal, flipped order",
			args: args{
				a: line{
					start:       coordinate{2, 5},
					end:         coordinate{7, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{11, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},

		{
			name: "lines partial overlap horizontal, good order",
			args: args{
				a: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{7, 5},
					end:         coordinate{16, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{16, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},
		{
			name: "lines partial overlap horizontal, good order",
			args: args{
				a: line{
					start:       coordinate{7, 5},
					end:         coordinate{16, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
				b: line{
					start:       coordinate{0, 5},
					end:         coordinate{11, 5},
					orientation: lineHorizontal,
					rowCol:      5,
				},
			},
			want: line{
				start:       coordinate{0, 5},
				end:         coordinate{16, 5},
				orientation: lineHorizontal,
				rowCol:      5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeLines(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("mergeLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pluck(t *testing.T) {
	type args[T any] struct {
		sl  []T
		idx int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    T
		want1   []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "fails to pluck int from slice",
			args: args[int]{
				sl:  []int{0, 1, 2, 3, 4, 5, 6},
				idx: 7,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "plucks int from slice",
			args: args[int]{
				sl:  []int{0, 1, 2, 3, 4, 5, 6},
				idx: 3,
			},
			want:    3,
			want1:   []int{0, 1, 2, 4, 5, 6},
			wantErr: false,
		},
		{
			name: "plucks int from end of slice",
			args: args[int]{
				sl:  []int{0, 1, 2, 3, 4, 5, 6},
				idx: 6,
			},
			want:    6,
			want1:   []int{0, 1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "plucks int from beginning of slice",
			args: args[int]{
				sl:  []int{0, 1, 2, 3, 4, 5, 6},
				idx: 0,
			},
			want:    0,
			want1:   []int{1, 2, 3, 4, 5, 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := pluck(tt.args.sl, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("pluck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pluck() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("pluck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reduceLines(t *testing.T) {
	type args struct {
		lines []line
	}
	tests := []struct {
		name    string
		args    args
		want    []line
		wantErr bool
	}{
		{
			name: "merges the two lines together",
			args: args{
				lines: []line{
					{
						start:       coordinate{0, 6},
						end:         coordinate{10, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{4, 6},
						end:         coordinate{12, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
				},
			},
			want: []line{
				{
					start:       coordinate{0, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
			},
			wantErr: false,
		},
		{
			name: "merges the two of the three lines together",
			args: args{
				lines: []line{
					{
						start:       coordinate{-10, 6},
						end:         coordinate{-5, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{0, 6},
						end:         coordinate{10, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{4, 6},
						end:         coordinate{12, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
				},
			},
			want: []line{
				{
					start:       coordinate{-10, 6},
					end:         coordinate{-5, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
				{
					start:       coordinate{0, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
			},
			wantErr: false,
		},
		{
			name: "merges the two of the three lines together, weird order",
			args: args{
				lines: []line{
					{
						start:       coordinate{4, 6},
						end:         coordinate{12, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{-10, 6},
						end:         coordinate{-5, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{0, 6},
						end:         coordinate{10, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
				},
			},
			want: []line{
				{
					start:       coordinate{-10, 6},
					end:         coordinate{-5, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
				{
					start:       coordinate{0, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
			},
			wantErr: false,
		},
		{
			name: "merges the three of the three lines together, weird order",
			args: args{
				lines: []line{
					{
						start:       coordinate{4, 6},
						end:         coordinate{12, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{-10, 6},
						end:         coordinate{4, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{0, 6},
						end:         coordinate{10, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
				},
			},
			want: []line{
				{
					start:       coordinate{-10, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
			},
			wantErr: false,
		},

		{
			name: "merges two groups together, weird order",
			args: args{
				lines: []line{
					{
						start:       coordinate{4, 6},
						end:         coordinate{12, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{-10, 6},
						end:         coordinate{4, 6},
						orientation: lineHorizontal,
						rowCol:      6,
					},
					{
						start:       coordinate{3, 0},
						end:         coordinate{3, 6},
						orientation: lineVertical,
						rowCol:      3,
					},
					{
						start:       coordinate{3, 4},
						end:         coordinate{3, 10},
						orientation: lineVertical,
						rowCol:      3,
					},
				},
			},
			want: []line{
				{
					start:       coordinate{-10, 6},
					end:         coordinate{12, 6},
					orientation: lineHorizontal,
					rowCol:      6,
				},
				{
					start:       coordinate{3, 0},
					end:         coordinate{3, 10},
					orientation: lineVertical,
					rowCol:      3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reduceLines(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("reduceLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduceLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_line_Len(t *testing.T) {
	type fields struct {
		start coordinate
		end   coordinate
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "point",
			fields: fields{
				start: coordinate{0, 0},
				end:   coordinate{0, 0},
			},
			want: 1,
		},
		{name: "4 long",
			fields: fields{
				start: coordinate{0, 4},
				end:   coordinate{0, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := newLine(tt.fields.start, tt.fields.end)
			if err != nil {
				t.Errorf("newline didn't work: %s", err.Error())
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
