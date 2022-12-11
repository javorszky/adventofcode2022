package day11

import (
	"reflect"
	"testing"
)

func Test_coolDownP1(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "6 divided by 3 is 2",
			args: args{in: 6},
			want: 2,
		},
		{
			name: "5 divided by 3 is 1",
			args: args{in: 5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coolDownP1(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("coolDownP1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateDivisibleFn(t *testing.T) {
	type args struct {
		divisor int
		num     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "40 divisible by 10",
			args: args{
				divisor: 10,
				num:     40,
			},
			want: true,
		},
		{
			name: "40 divisible by 11",
			args: args{
				divisor: 11,
				num:     40,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := generateModuloFn(tt.args.divisor)
			got := fn(tt.args.num)
			if got == 0 != tt.want {
				t.Errorf("generateModuloFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
