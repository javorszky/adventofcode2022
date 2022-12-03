package day3

import (
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:    "errors out for odd length",
			args:    args{in: "abc"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "errors out for empty input",
			args:    args{in: ""},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "splits string in two correctly",
			args:    args{in: "abcdefgh"},
			want:    "abcd",
			want1:   "efgh",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := split(tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("split() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("split() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_convertToPriority(t *testing.T) {
	type args struct {
		in int32
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "lowercase a is good priority",
			args:    args{in: 97},
			want:    1,
			wantErr: false,
		},
		{
			name:    "lowercase c is good priority",
			args:    args{in: 99},
			want:    3,
			wantErr: false,
		},
		{
			name:    "lowercase z is good priority",
			args:    args{in: 122},
			want:    26,
			wantErr: false,
		},
		{
			name:    "uppercase A is good priority",
			args:    args{in: 65},
			want:    27,
			wantErr: false,
		},
		{
			name:    "uppercase C is good priority",
			args:    args{in: 67},
			want:    29,
			wantErr: false,
		},
		{
			name:    "uppercase Z is good priority",
			args:    args{in: 90},
			want:    52,
			wantErr: false,
		},
		{
			name:    "too low codepoint results in error",
			args:    args{in: 64},
			want:    0,
			wantErr: true,
		},
		{
			name:    "between the ranges codepoint results in error",
			args:    args{in: 91},
			want:    0,
			wantErr: true,
		},
		{
			name:    "too high codepoint results in error",
			args:    args{in: 123},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToPriority(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToPriority() got = %v, want %v", got, tt.want)
			}
		})
	}
}
