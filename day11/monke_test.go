package day11

import "testing"

func Test_preSendExample(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remainder for 1",
			args: args{in: 1},
			want: 1,
		},
		{
			name: "remainder for the same number",
			args: args{in: examplePrimeProduct},
			want: 0,
		},
		{
			name: "remainder for some random number",
			args: args{in: 98843},
			want: 2266,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preSendExample(tt.args.in); got != tt.want {
				t.Errorf("preSendExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
