package day13

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    list
		want1   int
		wantErr bool
	}{
		{
			name:    "parses a single number",
			args:    args{line: "[1]"},
			want:    list{integer(1)},
			want1:   3,
			wantErr: false,
		},
		{
			name:    "parses an empty list",
			args:    args{line: "[]"},
			want:    list{},
			want1:   2,
			wantErr: false,
		},
		{
			name: "parses a single number, and a list",
			args: args{line: "[1,[2]]"},
			want: list{
				integer(1),
				list{
					integer(2),
				},
			},
			want1:   7,
			wantErr: false,
		},
		{
			name: "parses a list, and a single number",
			args: args{line: "[[2],1]"},
			want: list{
				list{
					integer(2),
				},
				integer(1),
			},
			want1:   7,
			wantErr: false,
		},
		{
			name: "parses a list-list, and a single number",
			args: args{line: "[[[1]],2]"},
			want: list{
				list{
					list{
						integer(1),
					},
				},
				integer(2),
			},
			want1:   9,
			wantErr: false,
		},
		{
			name: "parses a complex list",
			args: args{line: "[[[3,4],2,[[[5]]]],[5,2,[[2],8]]]"},
			want: list{
				list{
					list{
						integer(3),
						integer(4),
					},
					integer(2),
					list{
						list{
							list{
								integer(5),
							},
						},
					},
				},
				list{
					integer(5),
					integer(2),
					list{
						list{
							integer(2),
						},
						integer(8),
					},
				},
			},
			want1:   33,
			wantErr: false,
		},
		{
			name: "does some weird things",
			args: args{line: "[1234,[5678],90]"},
			want: list{
				integer(1234),
				list{
					integer(5678),
				},
				integer(90),
			},
			want1:   16,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseLine(tt.args.line, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseLine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_smallerList(t *testing.T) {
	type args struct {
		left  list
		right list
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "integers correct order",
			args: args{
				left: list{
					integer(3),
				},
				right: list{
					integer(4),
				},
			},
			want: correctOrder,
		},
		{
			name: "integers incorrect order",
			args: args{
				left: list{
					integer(3),
				},
				right: list{
					integer(2),
				},
			},
			want: incorrectOrder,
		},
		{
			name: "integers continue",
			args: args{
				left: list{
					integer(3),
				},
				right: list{
					integer(3),
				},
			},
			want: continueEvaluation,
		},
		{
			name: "shorter left list correct",
			args: args{
				left: list{
					list{
						integer(2),
					},
				},
				right: list{
					list{
						integer(2),
						integer(5),
					},
				},
			},
			want: correctOrder,
		},
		{
			name: "shorter right list incorrect",
			args: args{
				left: list{
					list{
						integer(2),
						integer(5),
						integer(5),
					},
				},
				right: list{
					list{
						integer(2),
						integer(5),
					},
				},
			},
			want: incorrectOrder,
		},
		{
			name: "same len list continue",
			args: args{
				left: list{
					list{
						integer(2),
						integer(5),
					},
				},
				right: list{
					list{
						integer(2),
						integer(5),
					},
				},
			},
			want: continueEvaluation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerList(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("smallerList() = %v, want %v", got, tt.want)
			}
		})
	}
}
